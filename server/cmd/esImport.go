package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8/typedapi/core/bulk"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/refresh"
	"github.com/spf13/cobra"
	"os"
	"server/global"
	"server/model/elasticsearch"
	"server/model/other"
	"server/service"
)

// esImportCmd represents the esImport command
var esImportCmd = &cobra.Command{
	Use:   "esImport",
	Short: "从指定的 JSON 文件导入数据到 ES",
	Long:  `从指定的 JSON 文件导入数据到 ES`,
	Run: func(cmd *cobra.Command, args []string) {
		elasticsearchImport, err := ElasticsearchImport("")
		if err != nil {
			panic(err)
		}
		fmt.Println("导入数据:", elasticsearchImport, "条")
	},
}

func init() {
	rootCmd.AddCommand(esImportCmd)
}

func ElasticsearchImport(jsonPath string) (int, error) {
	// 读取指定路径的 JSON 文件
	byteData, err := os.ReadFile(jsonPath)
	if err != nil {
		return 0, err
	}

	// 反序列化 JSON 数据到 ESIndexResponse 结构体
	var response other.ESIndexResponse
	err = json.Unmarshal(byteData, &response)
	if err != nil {
		return 0, err
	}

	// 创建 Elasticsearch 索引
	esService := service.ServiceGroupApp.EsService
	indexExists, err := esService.IndexExists(elasticsearch.ArticleIndex())
	if err != nil {
		return 0, err
	}
	if indexExists {
		if err := esService.IndexDelete(elasticsearch.ArticleIndex()); err != nil {
			return 0, err
		}
	}
	err = esService.IndexCreate(elasticsearch.ArticleIndex(), elasticsearch.ArticleMapping())
	if err != nil {
		return 0, err
	}

	// 构建批量请求数据
	var request bulk.Request
	for _, data := range response.Data {
		// 为每条数据创建索引操作，指定文档的 ID
		request = append(request, types.OperationContainer{Index: &types.IndexOperation{Id_: data.ID}})
		// 添加文档数据到请求
		request = append(request, data.Doc)
	}

	// 使用 Elasticsearch 客户端执行批量操作
	_, err = global.ESClient.Bulk().
		Request(&request).                   // 提交请求数据
		Index(elasticsearch.ArticleIndex()). // 指定索引名称
		Refresh(refresh.True).               // 强制刷新索引以使文档立即可见
		Do(context.TODO())                   // 执行请求
	if err != nil {
		return 0, err
	}

	// 返回导入的数据总条数
	total := len(response.Data)
	return total, nil
}
