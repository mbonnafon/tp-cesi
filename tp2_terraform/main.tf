resource "github_repository" "tp_cesi" {
  name        = "tp-cesi"
  description = "This is a super exercice !"

  visibility    = "public"
  has_downloads = true
  has_issues    = true
  has_projects  = true
  has_wiki      = true

  allow_auto_merge   = true
  allow_merge_commit = false
  allow_rebase_merge = false
  allow_squash_merge = false
}