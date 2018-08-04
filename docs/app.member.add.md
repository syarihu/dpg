NAME:
   dpg app member add - Invite users to the specified application

USAGE:
   dpg app member add [command options] [arguments...]

OPTIONS:
   --token value      [Required] API token
   --app-owner value  [Required] The owner of the application
   --app-id value     [Required] The application id. e.g. com.deploygate.sample
   --android          [Required] Specify this if the application is an android application
   --ios              [Required] Specify this if the application is an iOS application
   --invitees value   [Required] Comma-separated names or e-mails of those whom you want to invite
   