resource "sdwan_cedge_pim_feature_template" "example" {
  name         = "Example"
  description  = "My Example"
  device_types = ["vedge-C8000V"]
  auto_rp      = true
  rp_announce_fields = [
    {
      interface_name = "example"
      scope          = 1
    }
  ]
  interface_name = "example"
  rp_candidates = [
    {
      interface   = "example-interface"
      access_list = "1"
      interval    = 100
      priority    = 2
    }
  ]
  bsr_candidate            = "example-interface"
  hask_mask_length         = "100"
  priority                 = 1
  rp_candidate_access_list = "120"
  scope                    = 1
  range                    = "16"
  default                  = true
  rp_addresses = [
    {
      ip_address  = "1.2.3.4"
      access_list = "99"
      override    = false
    }
  ]
  spt_threshold = "0"
  interfaces = [
    {
      interface_name      = "example-interface"
      query_interval      = 30
      join_prune_interval = 60
    }
  ]
}
