output "app_ip" {
  description = "The public IP address of the app"
  value = aws_eip.app_eip[0].public_ip
  depends_on = [ aws_eip.app_eip ]
}

output "web_public_dns" {
  description = "The public DNS address of the app"
  value = aws_eip.app_eip[0].public_dns
  depends_on = [ aws_eip.app_eip ]
}

output "database_endpoint" {
  description = "The endpoint of the database"
  value = aws_db_instance.postgres.endpoint
}

output "database_port" {
  description = "The port of the database"
  value = aws_db_instance.postgres.port
}

output "redis_endpoint" {
  description = "The endpoint of redis"
  value = aws_elasticache_cluster.redis.cache_nodes[0].address
}

output "redis_port" {
  description = "The port of redis"
  value = aws_elasticache_cluster.redis.cache_nodes[0].port
}