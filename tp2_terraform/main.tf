resource "github_repository" "tp_cesi" {
  name        = "tp-cesi"

  visibility = "public"
  has_downloads = true
  has_issues = true
  has_projects = true
  has_wiki = true

}