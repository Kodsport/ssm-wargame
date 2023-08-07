package goa

import . "goa.design/goa/v3/dsl"

var _ = Service("admin", func() {
	Error("unauthorized")
	Error("not_found")
	Error("bad_request")
	HTTP(func() {
		Response("unauthorized", StatusForbidden, func() {
			Description("When the user does not have the required role")
		})
		Response("not_found", StatusNotFound, func() {
			Description("Resource not found")
		})
		Response("bad_request", StatusBadRequest, func() {
			Description("Resource not found")
		})

		Path("/admin")
	})
	Security(JWTAuth)

	Method("ListChallenges", func() {
		Payload(func() {
			Extend(TokenPayload)
		})
		Result(CollectionOf(ResultAdminChallenge))
		HTTP(func() {
			GET("/challenges")
			Response(StatusOK)
		})
	})

	Method("GetChallengeMeta", func() {
		Payload(func() {
			Extend(TokenPayload)
			Extend(ChallengeIDArtifact)
		})
		Result(ChallengeMeta)
		HTTP(func() {
			GET("/challenges/{challengeID}")
			Response(StatusOK)
		})
	})

	Method("CreateChallenge", func() {
		Payload(func() {
			Extend(CreateChallengePayload)
			Extend(TokenPayload)
		})
		HTTP(func() {
			POST("/challenges")
			Response(StatusCreated)
		})
	})

	Method("UpdateChallenge", func() {
		Payload(func() {
			Extend(CreateChallengePayload)
			Extend(TokenPayload)
			Extend(ChallengeIDArtifact)
		})
		HTTP(func() {
			PUT("/challenges/{challengeID}")
			Response(StatusCreated)
		})
	})

	Method("PresignChallFileUpload", func() {
		Payload(func() {
			Extend(TokenPayload)
			Extend(ChallengeIDArtifact)
			Attribute("md5", String, "MD5 hash of the file content in base64", func() {
				Example("cq02dBbcuugBHM1oKyvMlQ==")
			})
			Attribute("filename", String, func() {
				Example("decryptor.exe")
			})
			Attribute("size", Int64, func() {
				Example(41239)
				Description("the files number of bytes")
			})
			Required("md5", "filename", "size")
		})
		Result(func() {
			Attribute("url", String, "Signed PutObject URL ", func() {
				Example("https://s3-endpoint.example/bucket/key?signature=xxx")
			})
			Required("url")
		})
		HTTP(func() {
			POST("/challenges/{challengeID}/file_url")
			Response(StatusOK)
		})
	})

	Method("ListMonthlyChallenges", func() {
		Payload(func() {
			Extend(TokenPayload)
		})
		Result(ArrayOf(MonthlyChallenge))
		HTTP(func() {
			GET("/monthly_challenges")
		})
	})

	Method("DeleteMonthlyChallenge", func() {
		Payload(func() {
			Extend(TokenPayload)
			Extend(ChallengeIDArtifact)
		})
		HTTP(func() {
			DELETE("/monthly_challenges/{challengeID}")
		})
	})

	Method("DeleteFile", func() {
		Payload(func() {
			Extend(TokenPayload)
			Extend(FileIDArtifact)
		})
		HTTP(func() {
			DELETE("/files/{fileID}")
		})
	})

	Method("CreateMonthlyChallenge", func() {
		Payload(func() {
			Extend(TokenPayload)
			Extend(MonthlyChallenge)
		})
		HTTP(func() {
			POST("/monthly_challenges")
		})
	})

	Method("ListUsers", func() {
		Payload(func() {
			Extend(TokenPayload)
		})
		Result(ArrayOf(User))
		HTTP(func() {
			GET("/users")
		})
	})

	Method("AddFlag", func() {
		Payload(func() {
			Extend(TokenPayload)
			Attribute("flag", String, func() {
				Example("SSM{...}")
			})
			Extend(ChallengeIDArtifact)
			Required("flag")
		})

		HTTP(func() {
			POST("/challenges/{challengeID}/flags")
		})
	})

	Method("DeleteFlag", func() {
		Payload(func() {
			Extend(TokenPayload)
			Attribute("flagID", String, func() {
				Example("ac1c4362-c121-45a3-9745-8f8160a55f96")
				Format(FormatUUID)
			})
			Extend(ChallengeIDArtifact)
			Required("flagID")
		})

		HTTP(func() {
			DELETE("/challenges/{challengeID}/flags/{flagID}")
		})
	})

	Method("ListCategories", func() {
		Result(ArrayOf(Category))
		Payload(func() {
			Extend(TokenPayload)
		})
		HTTP(func() {
			GET("/categories")
		})
	})

	Method("ChalltoolsImport", func() {
		Payload(func() {
			Extend(ChallImport)
			Attribute("import_token", String, func() {
				Example("ctfimp_7ad44accdcca4c5ea10ea7ea61bec01b_efc6066f6ca0c6cd")
			})
			Required("import_token")
		})
		NoSecurity()
		HTTP(func() {
			POST("/push_challenge")
			Header("import_token:X-API-Key", String, "Auth token")
			Response(StatusOK)
		})
	})

	Method("ListCTFEvents", func() {
		Payload(func() {
			Extend(TokenPayload)
		})
		Result(ArrayOf(CTFEvent))
		HTTP(func() {
			GET("/events")
		})
	})

	Method("CreateCTFEvent", func() {
		Payload(func() {
			Extend(TokenPayload)
			Reference(CTFEvent)
			Attribute("name")
		})
		HTTP(func() {
			POST("/events")
		})
	})

	Method("DeleteCTFEvent", func() {
		Payload(func() {
			Extend(TokenPayload)
			Reference(CTFEvent)
			Attribute("id")
		})
		HTTP(func() {
			DELETE("/events/{id}")
		})
	})

	Method("CreateCTFEventImportToken", func() {
		Payload(func() {
			Extend(TokenPayload)
			Attribute("event_id", String, func() {
				Example("e3bb4dc5-9479-42ce-aed3-b41e8139fccb")
				Format(FormatUUID)
			})
		})
		Result(func() {
			Attribute("token")
			Required("token")
		})
		HTTP(func() {
			POST("/import_token")
		})
	})
})
