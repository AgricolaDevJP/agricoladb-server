data "github_release" "agricoladb_server" {
  owner       = "AgricolaDevJP"
  repository  = "agricoladb-server"
  retrieve_by = "latest"
}

locals {
  archive_filename = "agricoladb-server-lambda_${var.lambda_architecture}.zip"
  archive_download_url = [
    for asset in data.github_release.agricoladb_server.assets :
    asset if lookup(asset, "name") == local.archive_filename
  ][0].browser_download_url
}

resource "terraform_data" "server_lambda_archive" {
  triggers_replace = timestamp()

  provisioner "local-exec" {
    command = "wget \"$download_url\" -nv -O agricoladb-server-lambda.zip"
    environment = {
      download_url = local.archive_download_url
    }
  }
}
