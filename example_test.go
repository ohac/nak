package main

func ExampleEventBasic() {
	app.Run([]string{"nak", "event", "--ts", "1699485669"})
	// Output:
	// {"id":"36d88cf5fcc449f2390a424907023eda7a74278120eebab8d02797cd92e7e29c","pubkey":"79be667ef9dcbbac55a06295ce870b07029bfcdb2dce28d959f2815b16f81798","created_at":1699485669,"kind":1,"tags":[],"content":"hello from the nostr army knife","sig":"68e71a192e8abcf8582a222434ac823ecc50607450ebe8cc4c145eb047794cc382dc3f888ce879d2f404f5ba6085a47601360a0fa2dd4b50d317bd0c6197c2c2"}
}

func ExampleEventComplex() {
	app.Run([]string{"nak", "event", "--ts", "1699485669", "-k", "11", "-c", "skjdbaskd", "--sec", "17", "-t", "t=spam", "-e", "36d88cf5fcc449f2390a424907023eda7a74278120eebab8d02797cd92e7e29c", "-t", "r=https://abc.def?name=foobar;nothing"})
	// Output:
	// {"id":"19aba166dcf354bf5ef64f4afe69ada1eb851495001ee05e07d393ee8c8ea179","pubkey":"2fa2104d6b38d11b0230010559879124e42ab8dfeff5ff29dc9cdadd4ecacc3f","created_at":1699485669,"kind":11,"tags":[["t","spam"],["r","https://abc.def?name=foobar","nothing"],["e","36d88cf5fcc449f2390a424907023eda7a74278120eebab8d02797cd92e7e29c"]],"content":"skjdbaskd","sig":"cf452def4a68341c897c3fc96fa34dc6895a5b8cc266d4c041bcdf758ec992ec5adb8b0179e98552aaaf9450526a26d7e62e413b15b1c57e0cfc8db6b29215d7"}
}

func ExampleReq() {
	app.Run([]string{"nak", "req", "-k", "1", "-l", "18", "-a", "2fa2104d6b38d11b0230010559879124e42ab8dfeff5ff29dc9cdadd4ecacc3f", "-e", "aec4de6d051a7c2b6ca2d087903d42051a31e07fb742f1240970084822de10a6"})
	// Output:
	// ["REQ","nak",{"kinds":[1],"authors":["2fa2104d6b38d11b0230010559879124e42ab8dfeff5ff29dc9cdadd4ecacc3f"],"limit":18,"#e":["aec4de6d051a7c2b6ca2d087903d42051a31e07fb742f1240970084822de10a6"]}]
}

func ExampleEncodeNpub() {
	app.Run([]string{"nak", "encode", "npub", "a6a67ef9dcbbac55a06295ce870b07029bfcdb2dce28d959f2815b16f8179822"})
	// Output:
	// npub156n8a7wuhwk9tgrzjh8gwzc8q2dlekedec5djk0js9d3d7qhnq3qjpdq28
}

func ExampleEncodeNprofile() {
	app.Run([]string{"nak", "encode", "nprofile", "-r", "wss://example.com", "a6a67ef9dcbbac55a06295ce870b07029bfcdb2dce28d959f2815b16f8179822"})
	// Output:
	// nprofile1qqs2dfn7l8wthtz45p3ftn58pvrs9xlumvkuu2xet8egzkcklqtesgspz9mhxue69uhk27rpd4cxcefwvdhk6fl5jug
}
