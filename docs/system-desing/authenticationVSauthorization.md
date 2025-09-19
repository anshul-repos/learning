Authentication vs Authorization:

Authentication (Who You Are) : 

    • verifies identity
    • Analogy: At an airport, the security officer checks your passport or ID to verify that you are who you say you are.
    • Example: Logging into a system with a username/password.
    • OAuth, JWT, or session-based logins
    
Authorization (What You Can Do) :

    • what you’re allowed to do
    • Analogy: After verifying your ID, the boarding pass specifies which flight and seat you’re allowed to access.
    • Example: Once logged in, your role (e.g., admin/user) determines whether you can access certain files or features.
    • Role-Based Access Control (RBAC)


Easy Trick to Remember:

🔑 Authentication = "Can you prove who you are?"

🚦 Authorization = "Do you have permission?"


    Authentication  ->    Identification.  -> Verifies the identity.    -> e.g., username + password, fingerprint

    Authorization    ->    Permission.      -> verifies the access permission -> e.g., read/write access

