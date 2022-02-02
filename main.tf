#Data block
data "aws_route53_zone" "dns" {
  name         = "devops.rebrain.srwx.net."
  private_zone = false
}
data "digitalocean_ssh_key" "another" {
  name = "another.SSH.PUB.KEY"
}
#create resource my_ssh_key
resource "digitalocean_ssh_key" "my_ssh_dkocheto1" {
  name       = "my_ssh_dkocheto1"
  public_key = file("/home/inblade/.ssh/id_rsa.pub")
}
#create tags for DO droplet
resource "digitalocean_tag" "devops" {
  name = "devops"
}
resource "digitalocean_tag" "mail" {
  name = "dkocheto_at_gmail_com"
}
#create DO droplet
resource "digitalocean_droplet" "Dkocheto" {
  image    = "ubuntu-20-04-x64"
  name     = "Dkocheto"
  region   = "nyc3"
  size     = "s-1vcpu-1gb"
  tags     = [digitalocean_tag.devops.id, digitalocean_tag.mail.id]
  ssh_keys = [digitalocean_ssh_key.my_ssh_dkocheto1.id, data.digitalocean_ssh_key.another.id]

}

resource "aws_route53_record" "dns_rebrain" {
  zone_id = data.aws_route53_zone.dns.zone_id
  name    = var.devs
  type    = "A"
  ttl     = "30"
  records = [digitalocean_droplet.Dkocheto.ipv4_address]
}
/*resource "null_resource" "maingo" {
  provisioner "local-exec" {
    command = "cp ./main.go /tmp/"
  }
  depends_on = [aws_route53_record.dns_rebrain]
}*/

resource "local_file" "hosts" {
  filename = "hosts"
  content = templatefile("${path.module}/inventory.tpl", {
    ip = digitalocean_droplet.Dkocheto.*.ipv4_address
  })
}
