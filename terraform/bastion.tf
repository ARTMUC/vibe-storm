# Key Pair for Bastion Host
resource "aws_key_pair" "bastion" {
  key_name   = "${var.project_name}-${var.environment}-bastion-key"
  public_key = file("${path.module}/bastion.pub")
}

# Bastion Host EC2 Instance
resource "aws_instance" "bastion" {
  ami                         = "ami-0c02fb55956c7d316" # Amazon Linux 2 AMI (change as needed for your region)
  instance_type               = var.bastion_instance_type
  key_name                    = aws_key_pair.bastion.key_name
  vpc_security_group_ids      = [aws_security_group.bastion.id]
  subnet_id                   = aws_subnet.public[0].id
  associate_public_ip_address = true

  tags = {
    Name = "${var.project_name}-${var.environment}-bastion"
  }
}
