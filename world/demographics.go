package world

import "github.com/cjduffett/synthea/utils"

// Demographics is a lookup hash for demographics used to seed Synthea.
var Demographics = map[string]interface{}{

	"Race": []utils.Choice{
		// https://en.wikipedia.org/wiki/Demographics_of_Massachusetts#Race.2C_ethnicity.2C_and_ancestry
		utils.Choice{
			Weight: 0.694,
			Item:   "White",
		},
		utils.Choice{
			Weight: 0.105,
			Item:   "Hispanic",
		},
		utils.Choice{
			Weight: 0.081,
			Item:   "Black",
		},
		utils.Choice{
			Weight: 0.06,
			Item:   "Asian",
		},
		utils.Choice{
			Weight: 0.05,
			Item:   "Native",
		},
		utils.Choice{
			Weight: 0.01,
			Item:   "Other",
		},
	},

	"Ethnicity": map[string][]utils.Choice{
		"White": []utils.Choice{
			// https://en.wikipedia.org/wiki/Demographics_of_Massachusetts#Race.2C_ethnicity.2C_and_ancestry
			// Scaled out of 100% for each ethnicity.
			utils.Choice{
				Weight: 0.263,
				Item:   "Irish",
			},
			utils.Choice{
				Weight: 0.160,
				Item:   "Italian",
			},
			utils.Choice{
				Weight: 0.123,
				Item:   "English",
			},
			utils.Choice{
				Weight: 0.090,
				Item:   "French",
			},
			utils.Choice{
				Weight: 0.074,
				Item:   "German",
			},
			utils.Choice{
				Weight: 0.057,
				Item:   "Polish",
			},
			utils.Choice{
				Weight: 0.054,
				Item:   "Portuguese",
			},
			utils.Choice{
				Weight: 0.050,
				Item:   "American",
			},
			utils.Choice{
				Weight: 0.044,
				Item:   "French Canadian",
			},
			utils.Choice{
				Weight: 0.028,
				Item:   "Scottish",
			},
			utils.Choice{
				Weight: 0.022,
				Item:   "Russian",
			},
			utils.Choice{
				Weight: 0.021,
				Item:   "Swedish",
			},
			utils.Choice{
				Weight: 0.014,
				Item:   "Greek",
			},
		},
		"Hispanic": []utils.Choice{
			utils.Choice{
				Weight: 0.577,
				Item:   "Puerto Rican",
			},
			utils.Choice{
				Weight: 0.0141,
				Item:   "Mexican",
			},
			utils.Choice{
				Weight: 0.0141,
				Item:   "Central American",
			},
			utils.Choice{
				Weight: 0.0141,
				Item:   "South American",
			},
		},
		"Black": []utils.Choice{
			utils.Choice{
				Weight: 0.34,
				Item:   "African",
			},
			utils.Choice{
				Weight: 0.33,
				Item:   "Dominican",
			},
			utils.Choice{
				Weight: 0.33,
				Item:   "West Indian",
			},
		},
		"Asian": []utils.Choice{
			utils.Choice{
				Weight: 0.6,
				Item:   "Chinese",
			},
			utils.Choice{
				Weight: 0.4,
				Item:   "Asian Indian",
			},
		},
		"Native": []utils.Choice{
			utils.Choice{
				Weight: 1,
				Item:   "American Indian",
			},
		},
		"Other": []utils.Choice{
			utils.Choice{
				Weight: 1,
				Item:   "Arab",
			},
		},
	},

	"BloodType": map[string][]utils.Choice{
		// blood type data from http://www.redcrossblood.org/learn-about-blood/blood-types
		// data for Native and Other from https://en.wikipedia.org/wiki/Blood_type_distribution_by_country
		"White": []utils.Choice{
			utils.Choice{
				Weight: 0.37,
				Item:   "o_positive",
			},
			utils.Choice{
				Weight: 0.08,
				Item:   "o_negative",
			},
			utils.Choice{
				Weight: 0.33,
				Item:   "a_positive",
			},
			utils.Choice{
				Weight: 0.07,
				Item:   "a_negative",
			},
			utils.Choice{
				Weight: 0.09,
				Item:   "b_positive",
			},
			utils.Choice{
				Weight: 0.02,
				Item:   "b_negative",
			},
			utils.Choice{
				Weight: 0.03,
				Item:   "ab_positive",
			},
			utils.Choice{
				Weight: 0.01,
				Item:   "ab_negative",
			},
		},
		"Hispanic": []utils.Choice{
			utils.Choice{
				Weight: 0.52,
				Item:   "o_positive",
			},
			utils.Choice{
				Weight: 0.04,
				Item:   "o_negative",
			},
			utils.Choice{
				Weight: 0.29,
				Item:   "a_positive",
			},
			utils.Choice{
				Weight: 0.02,
				Item:   "a_negative",
			},
			utils.Choice{
				Weight: 0.09,
				Item:   "b_positive",
			},
			utils.Choice{
				Weight: 0.01,
				Item:   "b_negative",
			},
			utils.Choice{
				Weight: 0.02,
				Item:   "ab_positive",
			},
			utils.Choice{
				Weight: 0.01,
				Item:   "ab_negative",
			},
		},
		"Black": []utils.Choice{
			utils.Choice{
				Weight: 0.46,
				Item:   "o_positive",
			},
			utils.Choice{
				Weight: 0.04,
				Item:   "o_negative",
			},
			utils.Choice{
				Weight: 0.24,
				Item:   "a_positive",
			},
			utils.Choice{
				Weight: 0.02,
				Item:   "a_negative",
			},
			utils.Choice{
				Weight: 0.18,
				Item:   "b_positive",
			},
			utils.Choice{
				Weight: 0.01,
				Item:   "b_negative",
			},
			utils.Choice{
				Weight: 0.04,
				Item:   "ab_positive",
			},
			utils.Choice{
				Weight: 0.01,
				Item:   "ab_negative",
			},
		},
		"Asian": []utils.Choice{
			utils.Choice{
				Weight: 0.39,
				Item:   "o_positive",
			},
			utils.Choice{
				Weight: 0.01,
				Item:   "o_negative",
			},
			utils.Choice{
				Weight: 0.26,
				Item:   "a_positive",
			},
			utils.Choice{
				Weight: 0.01,
				Item:   "a_negative",
			},
			utils.Choice{
				Weight: 0.25,
				Item:   "b_positive",
			},
			utils.Choice{
				Weight: 0.01,
				Item:   "b_negative",
			},
			utils.Choice{
				Weight: 0.06,
				Item:   "ab_positive",
			},
			utils.Choice{
				Weight: 0.01,
				Item:   "ab_negative",
			},
		},
		"Native": []utils.Choice{
			utils.Choice{
				Weight: 0.374,
				Item:   "o_positive",
			},
			utils.Choice{
				Weight: 0.066,
				Item:   "o_negative",
			},
			utils.Choice{
				Weight: 0.357,
				Item:   "a_positive",
			},
			utils.Choice{
				Weight: 0.063,
				Item:   "a_negative",
			},
			utils.Choice{
				Weight: 0.085,
				Item:   "b_positive",
			},
			utils.Choice{
				Weight: 0.015,
				Item:   "b_negative",
			},
			utils.Choice{
				Weight: 0.034,
				Item:   "ab_positive",
			},
			utils.Choice{
				Weight: 0.006,
				Item:   "ab_negative",
			},
		},
		"Other": []utils.Choice{
			utils.Choice{
				Weight: 0.374,
				Item:   "o_positive",
			},
			utils.Choice{
				Weight: 0.066,
				Item:   "o_negative",
			},
			utils.Choice{
				Weight: 0.357,
				Item:   "a_positive",
			},
			utils.Choice{
				Weight: 0.063,
				Item:   "a_negative",
			},
			utils.Choice{
				Weight: 0.085,
				Item:   "b_positive",
			},
			utils.Choice{
				Weight: 0.015,
				Item:   "b_negative",
			},
			utils.Choice{
				Weight: 0.034,
				Item:   "ab_positive",
			},
			utils.Choice{
				Weight: 0.006,
				Item:   "ab_negative",
			},
		},
	},
}
