resource "sdwan_cellular_profile_feature_template" "example" {
  name                     = "Example"
  description              = "My Example"
  device_types             = ["vedge-C8000V"]
  if_name                  = "example"
  profile_id               = 1
  access_point_name        = "example"
  authentication_type      = "CHAP"
  ip_address               = "1.2.3.4"
  profile_name             = "example"
  packet_data_network_type = "ipv4"
  profile_username         = "MyUsername"
  profile_password         = "MyPassword"
  primary_dns_address      = "1.2.3.4"
  secondary_dns_address    = "1.2.3.4"
}
