package books

const BOOK_ONE = `{
	"title": "Harry Potter and the Philosophers Stone",
	"author": "J.K. Rowling",
	"year": 1997,
	"publisher": "Bloomsbury (UK)",
	"description": "A book about a wizard boy"
	}`

const BOOK_ONE_RESPONSE = `{
		"id": 1
		}`
const BOOK_TWO = `{
	"title": "Old Testament",
	"author": "Various",
	"year": -165,
	"description": "A holy book of Christianity and Jewish faith"
	}`
const BOOK_TWO_RESPONSE = `{
	"id": 2
	}`

const BOOK_THREE = `{
		"title": "The Subtle Knife",
		"author": "Philip Pullman",
		"year": 1997,
		"publisher": "Scholastic Point"
		}`

const BOOK_THREE_RESPONSE = `{
		"id": 3
		}`

const BOOK_FOUR = `{
			"title": "Goosebumps: Beware, the Snowman",
			"author": "R.L. Stine",
			"year": 1997,
			"publisher": "Scholastic Point"
			}`

const BOOK_FOUR_RESPONSE = `{
			"id": 4
			}`

const BOOK_WITH_MISSING_TITLE = `{
				"author": "Douglas Adams",
				"year": 1979,
				"publisher": "Pan Books",
				"description": "Originally a radio series"
				}`

const BOOK_WITH_MISSING_YEAR = `{
					"title": "The Hitchhiker&#39;s Guide to the Galaxy",
					"author": "Douglas Adams",
					"publisher": "Pan Books",
					"description": "Originally a radio series"
					}`

const BOOK_WITH_MISSING_YEAR_TWO = `{
						"author": "Douglas Adams",
						"title": "The Hitchhiker&#39;s Guide to the Galaxy",
						"pages": 208,
						
						"description": "Originally a radio series"
						}`
const BOOK_WITH_EMPTY_AUTHOR = `{
						"author": "",
						"title": "The Hitchhiker&#39;s Guide to the Galaxy",
						"year": 1979,
						"publisher": "Pan Books",
						"description": "Originally a radio series"
						}`
const BOOK_WITH_EMPTY_TITLE = `{
						"author": "Douglas Adams",
						"title": "",
						"year": 1979,
						"publisher": "Pan Books",
						"description": "Originally a radio series"
						}`

const BOOK_WITH_NON_INTEGER_YEAR = `{
							"author": "Douglas Adams",
							"title": "",
							"year": 1979,999,
							"publisher": "Pan Books",
							"description": "Originally a radio series"
							}`

const BOOK_WITH_NON_INTEGER_YEAR_TWO = `{
								"author": "Douglas Adams",
								"title": "",
								"year": "nineteen-ninety-seven",
								"publisher": "Pan Books",
								"description": "Originally a radio series"
								}`

const BOOK_WITH_EMPTY_PUBLISHER = `{
									"author": "Douglas Adams",
									"title": "",
									
									"year": 1979,
									"publisher": "",
									"description": "Originally a radio series"
									}`

const BOOK_WITH_NON_UNIQUE_VALUES = `{
										"title": "Harry Potter and the Philosophers Stone",
										"author": "J.K. Rowling",
										"year": 1997,
										"publisher": "Bloomsbury (UK)",
										"description": "A book about a wizard boy"
										}`
const INVALID_JSON = `{
	"whoops"
	}`

const ALL_BOOKS_RESPONSE = `[
		{
		"id": 1,
		"title": "Harry Potter and the Philosophers Stone",
		"author": "J.K. Rowling",
		"year": 1997,
		"publisher": "Bloomsbury (UK)",
		"description": "A book about a wizard boy"
		},
		{
		"id": 2,
		"title": "Old Testament",
		"author": "Various",
		"year": -165,
		"publisher": null,
		"description": "A holy book of Christianity and Jewish faith"
		},
		{
		"id": 3,
		"title": "The Subtle Knife",
		"author": "Philip Pullman",
		"year": 1997,
		"publisher": "Scholastic Point",
		"description": null
		},
		{
		"id": 4,
		"title": "Goosebumps: Beware, the Snowman",
		"author": "R.L. Stine",
		"year": 1997,
		"publisher": "Scholastic Point",
		"description": null
		}
		]`

const BOOKS_FILTERED_BY_AUTHOR_RESPONSE = `[
			{
			"id": 1,
			"title": "Harry Potter and the Philosophers Stone",
			"author": "J.K. Rowling",
			"year": 1997,
			"publisher": "Bloomsbury (UK)",
			"description": "A book about a wizard boy"
			}
			]`

const BOOKS_FILTERD_BY_YEAR_RESPONSE = `[
				{
				"id": 1,
				"title": "Harry Potter and the Philosophers Stone",
				"author": "J.K. Rowling",
				"year": 1997,
				"publisher": "Bloomsbury (UK)",
				"description": "A book about a wizard boy"
				},
				{
				"id": 3,
				"title": "The Subtle Knife",
				"author": "Philip Pullman",
				"year": 1997,
				"publisher": "Scholastic Point",
				"description": null
				},
				{
				"id": 4,
				
				"title": "Goosebumps: Beware, the Snowman",
				"author": "R.L. Stine",
				"year": 1997,
				"publisher": "Scholastic Point",
				"description": null
				}
				]`

const BOOKS_FILTERED_BY_NONEXISTING_PUBLISHER_RESPONSE = `[]`

const BOOKS_FILTERED_BY_YEAR_AND_PUBLISHER_RESPONSE = `[
	{
	"id": 3,
	"title": "The Subtle Knife",
	"author": "Philip Pullman",
	"year": 1997,
	"publisher": "Scholastic Point",
	"description": null
	},
	{
	"id": 4,
	"title": "Goosebumps: Beware, the Snowman",
	"author": "R.L. Stine",
	"year": 1997,
	"publisher": "Scholastic Point",
	"description": null
	
	}
	]`

const BOOK_ONE_GET_RESPONSE = `{"id": 1, "title": "Harry Potter and the Philosophers Stone","author": "J.K. Rowling","year": 1997,"publisher": "Bloomsbury (UK)","description": "A book about a wizard boy"}`

const BOOKS_FINAL_STATE = `[
	{
	"id": 2,
	"title": "Old Testament",
	"author": "Various",
	"year": -165,
	"publisher": null,
	"description": "A holy book of Christianity and Jewish faith"
	},
	{
	"id": 3,
	"title": "The Subtle Knife",
	"author": "Philip Pullman",
	"year": 1997,
	"publisher": "Scholastic Point",
	"description": null
	},
	{
	"id": 4,
	"title": "Goosebumps: Beware, the Snowman",
	"author": "R.L. Stine",
	"year": 1997,
	"publisher": "Scholastic Point",
	"description": null
	}
	]`
