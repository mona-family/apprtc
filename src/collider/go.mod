module mona.family/collider

go 1.15

require (
	cloud.google.com/go/firestore v1.4.0 // indirect
	firebase.google.com/go v3.13.0+incompatible
	firebase.google.com/go/v4 v4.2.0
	golang.org/x/net v0.0.0-20210220033124-5f55cee0dc0d
	google.golang.org/api v0.40.0
)

replace firebase.google.com/go/v4 => github.com/mona-family/firebase-admin-go/v4 v4.2.1
