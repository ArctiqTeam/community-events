resource "helm_release" "beer" {
  name          = "beer"
  namespace     = "beer"
  chart         = "beer/"
  recreate_pods = true

  values = [
    "name: beer",
    "image: arctiqteam/random-beer-selector",
    "image_version: latest",
  ]
}
