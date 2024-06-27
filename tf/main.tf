terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
    }
  }
}

provider "aws" {
  region = var.aws_region
}

# Availability zones
data "aws_availability_zones" "available" {
  state = "available"
}

# Create VPC
resource "aws_vpc" "main" {
  cidr_block = var.vpc_cidr_block
  enable_dns_hostnames = true
  tags = {
    Name = var.app_name
  }
}

# Create Internet Gateway
resource "aws_internet_gateway" "main" {
  vpc_id = aws_vpc.main.id
  tags = {
    Name = var.app_name
  }
}

# Create Subnet Public
resource "aws_subnet" "public" {
  count             = var.subnet_count.public
  vpc_id            = aws_vpc.main.id
  cidr_block        = var.public_subnet_cidr_blocks[count.index]
  availability_zone = data.aws_availability_zones.available.names[count.index]

  tags = {
    Name = "${var.app_name}"
  }
}

# Create Subnet Private
resource "aws_subnet" "private" {
  count             = var.subnet_count.private
  vpc_id            = aws_vpc.main.id
  cidr_block        = var.private_subnet_cidr_blocks[count.index]
  availability_zone = data.aws_availability_zones.available.names[count.index]
  tags = {
    Name = "${var.app_name}"
  }
}

# Create Route Table Public
resource "aws_route_table" "public" {
  vpc_id = aws_vpc.main.id

  route {
    cidr_block = "0.0.0.0/0"
    gateway_id = aws_internet_gateway.main.id
  }

  tags = {
    Name = "${var.app_name}"
  }
}

# Associate Route Table with Subnet Public
resource "aws_route_table_association" "public" {
  count          = var.subnet_count.public
  route_table_id = aws_route_table.public.id
  subnet_id      = aws_subnet.public[count.index].id
}

# Create Route Table Private
resource "aws_route_table" "private" {
  vpc_id = aws_vpc.main.id

  tags = {
    Name = "${var.app_name}"
  }
}

# Associate Route Table with Subnet Private
resource "aws_route_table_association" "private" {
  count          = var.subnet_count.private
  route_table_id = aws_route_table.private.id
  subnet_id      = aws_subnet.private[count.index].id
}

# Create DB Subnet Group
resource "aws_db_subnet_group" "main" {
  name       = "db-subnet-group"
  subnet_ids = [ for subnet in aws_subnet.private : subnet.id ]

  tags = {
    Name = "${var.app_name}"
  }
}

# Create Redis Subnet Group
resource "aws_elasticache_subnet_group" "main" {
  name       = "redis-subnet-group"
  subnet_ids = [ for subnet in aws_subnet.private : subnet.id ]

  tags = {
    Name = "${var.app_name}"
  }
}
