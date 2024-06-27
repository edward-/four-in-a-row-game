variable "aws_region" {
  description = "The AWS region to deploy"
  default     = "us-east-1"
}

variable "vpc_cidr_block" {
  description = "The CIDR block for the VPC"
  type        = string
  default     = "10.0.0.0/16"
}

variable "subnet_count" {
  description = "Number of subnet"
  type = map(number)
  default = {
    public = 1,
    private = 2
  }
}

variable "settings" {
  description = "Configuration settings"
  type = map(any)
  default = {
    "database" = {
      allocated_storage    = 20
      engine               = "postgres"
      engine_version       = "16"
      instance_class       = "db.t4g.micro"
      db_name              = "rowgame_postgres_db"
      parameter_group_name = "default.postgres16"
      skip_final_snapshot  = true
    },
    "redis" = {
      cluster_id           = "redis-cluster"
      engine               = "memcached"
      node_type            = "cache.t4g.micro"
      num_cache_nodes      = 1
      parameter_group_name = "default.memcached1.6"
      engine_version       = "1.6.22"
      port                 = 6379
    },
    "app" = {
      count         = 1
      instance_type = "t2.micro" // architecture requirement
      ami           = "ami-08a0d1e16fc3f61ea"
    }
  }
}

variable "public_subnet_cidr_blocks" {
  description = "Available CIDR blocks for public subnets"
  type = list(string)
  default = [
    "10.0.1.0/24",
    "10.0.2.0/24",
    "10.0.3.0/24",
    "10.0.4.0/24"
  ]
}

variable "private_subnet_cidr_blocks" {
  description = "Available CIDR blocks for private subnets"
  type = list(string)
  default = [
    "10.0.101.0/24",
    "10.0.102.0/24",
    "10.0.103.0/24",
    "10.0.104.0/24",
  ]
}

variable "my_ip" {
  description = "Your IP address"
  type = string
  sensitive = true
}

variable "db_name" {
  description = "The name of the PostgreSQL database"
  default     = "rowgame_postgres_db"
}

variable "db_username" {
  description = "The PostgreSQL username"
  type        = string
  sensitive   = true
}

variable "db_password" {
  description = "The PostgreSQL password"
  type        = string
  sensitive   = true
}

variable "app_name" {
  description = "Application name"
  default = "4InARowGameApp"
}

variable "app_environment" {
  description = "Environment"
  default = "dev" 
}
