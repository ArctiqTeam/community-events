      resource "//cloudresourcemanager.googleapis.com/folders/ADD_YOUR_FOLDER_ID_HERE_###" { 
        roles = ["roles/resourcemanager.folderViewer", 
                 "roles/resourcemanager.projectCreator", 
                 "roles/storage.admin", 
                 "roles/iam.serviceAccountAdmin",
                 "roles/viewer", 
                 "roles/editor", 
                 "roles/resourcemanager.projectDeleter"]
      }
      resource "//cloudresourcemanager.googleapis.com/organizations/ADD_YOUR_ORG_ID_HERE" {
        roles = ["roles/billing.user", "roles/resourcemanager.organizationViewer"]
      }