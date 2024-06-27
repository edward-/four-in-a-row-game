
# Create RDS instance
resource "aws_db_instance" "postgres" {
  allocated_storage    = var.settings.database.allocated_storage
  engine               = var.settings.database.engine
  engine_version       = var.settings.database.engine_version
  instance_class       = var.settings.database.instance_class
  db_name              = var.settings.database.db_name
  username             = var.db_username
  password             = var.db_password
  parameter_group_name = var.settings.database.parameter_group_name
  skip_final_snapshot  = var.settings.database.skip_final_snapshot
  db_subnet_group_name = aws_db_subnet_group.main.id
  vpc_security_group_ids = [aws_security_group.rds_sg.id]

  tags = {
    Name = "${var.app_name}"
  }
}

# Create ElasticCache Redis instance
resource "aws_elasticache_cluster" "redis" {
  cluster_id           = var.settings.redis.cluster_id
  engine               = var.settings.redis.engine
  engine_version       = var.settings.redis.engine_version
  node_type            = var.settings.redis.node_type
  num_cache_nodes      = var.settings.redis.num_cache_nodes
  parameter_group_name = var.settings.redis.parameter_group_name
  port                 = var.settings.redis.port
  security_group_ids   = [aws_security_group.redis_sg.id]
  subnet_group_name    = aws_elasticache_subnet_group.main.name

  tags = {
    Name = "${var.app_name}"
  }
}

# Create EC2 instance
resource "aws_instance" "app" {
  count           = var.settings.app.count
  ami             = var.settings.app.ami
  instance_type   = var.settings.app.instance_type
  subnet_id       = aws_subnet.public[count.index].id
  
  vpc_security_group_ids = [aws_security_group.ec2_sg.id]

  tags = {
    Name = "${var.app_name}"
  }
}

resource "aws_eip" "app_eip" {
  count = var.settings.app.count
  instance = aws_instance.app[count.index].id
  domain = "vpc"

  tags = {
    Name = "${var.app_name}"
  }
}