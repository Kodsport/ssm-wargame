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
})
