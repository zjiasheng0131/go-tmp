provider "aws" {
  region = "us-west-2" # 替换为你的 AWS 区域
}

resource "aws_elasticsearch_domain" "example" {
  domain_name           = "my-elasticsearch-domain"
  elasticsearch_version = "7.10" # 替换为你需要的 Elasticsearch 版本

  cluster_config {
    instance_type = "t3.small.elasticsearch" # 替换为你需要的实例类型
    instance_count = 1
  }

  ebs_options {
    ebs_enabled = true
    volume_size = 10 # 存储大小（GB）
  }

  # 关闭 HTTPS，仅允许 HTTP 访问
  domain_endpoint_options {
    enforce_https       = false
    tls_security_policy = "Policy-Min-TLS-1-0-2019-07"
  }

  # 允许公开访问（仅用于测试，生产环境请谨慎使用）
  access_policies = <<POLICY
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Principal": {
        "AWS": "*"
      },
      "Action": "es:*",
      "Resource": "arn:aws:es:us-west-2:123456789012:domain/my-elasticsearch-domain/*"
    }
  ]
}
POLICY

  # 配置日志记录（可选）
  log_publishing_options {
    cloudwatch_log_group_arn = aws_cloudwatch_log_group.example.arn
    log_type                 = "INDEX_SLOW_LOGS"
  }
}

# CloudWatch Log Group（可选）
resource "aws_cloudwatch_log_group" "example" {
  name = "/aws/elasticsearch/my-elasticsearch-domain"
}

output "elasticsearch_endpoint" {
  value = aws_elasticsearch_domain.example.endpoint
}
