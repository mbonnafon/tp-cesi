resource "github_repository" "tp_cesi" {
  name = "tp-cesi"

  visibility    = "public"
  has_downloads = true
  has_issues    = true
  has_projects  = true
  has_wiki      = true

  allow_merge_commit   = false
  allow_rebase_merge   = false
  allow_squash_merge   = false

  default_branch = "main"
  description    = "This is a sample repository"
  topics         = ["terraform", "github", "repository", "test"]
}