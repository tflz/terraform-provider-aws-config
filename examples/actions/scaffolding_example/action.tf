resource "terraform_data" "example" {
  input = "fake-string"

  lifecycle {
    action_trigger {
      events  = [before_create]
      actions = [action.aws-landing-zone-config_example.example]
    }
  }
}

action "aws-landing-zone-config_example" "example" {
  config {
    configurable_attribute = "some-value"
  }
}