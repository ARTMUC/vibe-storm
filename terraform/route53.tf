# Route53 Hosted Zone
resource "aws_route53_zone" "main" {
  name = var.hosted_zone_domain
}

# Route53 DNS Record for Application
resource "aws_route53_record" "app" {
  zone_id = aws_route53_zone.main.zone_id
  name    = var.domain_name
  type    = "A"

  alias {
    name                   = aws_lb.main.dns_name
    zone_id                = aws_lb.main.zone_id
    evaluate_target_health = true
  }
}

# Route53 DNS Record for ACM Certificate Validation
resource "aws_route53_record" "cert_validation" {
  for_each = {
    for dvo in aws_acm_certificate.main.domain_validation_options : dvo.domain_name => {
      name   = dvo.resource_record_name
      record = dvo.resource_record_value
      type   = dvo.resource_record_type
    }
  }

  allow_overwrite = true
  name            = each.value.name
  records         = [each.value.record]
  ttl             = 60
  type            = each.value.type
  zone_id         = aws_route53_zone.main.zone_id
}

# Route53 DNS Record for Database (private)
resource "aws_route53_record" "db" {
  zone_id = aws_route53_zone.main.zone_id
  name    = "db.${var.domain_name}"
  type    = "CNAME"
  ttl     = "300"
  records = [aws_db_instance.main.endpoint]
}
