# Security Group for EC2 instance
resource "aws_security_group" "ec2_sg" {
  name    = "${var.app_name}_ec2_sg"
  vpc_id = aws_vpc.main.id

  ingress {
    description = "Allow all traffic through HTTP"
    from_port   = 80
    to_port     = 80
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  ingress {
    description = "Allow SSH from my computer"
    from_port   = 22
    to_port     = 22
    protocol    = "tcp"
    cidr_blocks = ["${var.my_ip}/32"]
  }

  egress {
    description = "Allow all outbound traffic"
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags = {
    Name = "${var.app_name}"
  }
}

# Security Group for RDS instance
resource "aws_security_group" "rds_sg" {
  name = "${var.app_name}_rds_sg"
  vpc_id = aws_vpc.main.id

  ingress {
    description     = "Allow Postgres traffic from only ec2 sg"
    from_port       = 5432
    to_port         = 5432
    protocol        = "tcp"
    security_groups = [aws_security_group.ec2_sg.id]
  }

  tags = {
    Name = "${var.app_name}"
  }
}

# Security Group for ElasticCache Redis
resource "aws_security_group" "redis_sg" {
  name = "${var.app_name}_redis_sg"
  vpc_id = aws_vpc.main.id

  ingress {
    description     = "Allow Redis traffic from only ec2 sg"
    from_port       = 6379
    to_port         = 6379
    protocol        = "tcp"
    security_groups = [aws_security_group.ec2_sg.id]
  }

  tags = {
    Name = "${var.app_name}"
  }
}