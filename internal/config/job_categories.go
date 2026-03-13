package config

type JobCategory struct {
	Key     string `json:"key"`
	LabelJa string `json:"label_ja"`
	LabelEn string `json:"label_en"`
	Icon    string `json:"icon"`
}

var JobCategories = []JobCategory{
	{Key: "beauty",           LabelJa: "美容師",               LabelEn: "Beautician",        Icon: "💇"},
	{Key: "nail_art",         LabelJa: "ネイリスト",            LabelEn: "Nail Artist",        Icon: "💅"},
	{Key: "eyelash",          LabelJa: "まつ毛エクステ",        LabelEn: "Eyelash Extension",  Icon: "👁"},
	{Key: "massage",          LabelJa: "マッサージ",            LabelEn: "Massage Therapist",  Icon: "💆"},
	{Key: "facial",           LabelJa: "フェイシャルエステ",    LabelEn: "Facial Esthetician", Icon: "✨"},
	{Key: "hair_removal",     LabelJa: "脱毛",                 LabelEn: "Hair Removal",       Icon: "🌿"},
	{Key: "makeup",           LabelJa: "メイクアップ",          LabelEn: "Makeup Artist",      Icon: "💄"},
	{Key: "hair_stylist",     LabelJa: "ヘアスタイリスト",      LabelEn: "Hair Stylist",       Icon: "💈"},
	{Key: "barber",           LabelJa: "理容師",               LabelEn: "Barber",             Icon: "✂️"},
	{Key: "spa",              LabelJa: "スパセラピスト",        LabelEn: "Spa Therapist",      Icon: "🛁"},
	{Key: "waxing",           LabelJa: "ワックス脱毛",          LabelEn: "Waxing",             Icon: "🕯️"},
	{Key: "tattoo",           LabelJa: "タトゥーアーティスト",  LabelEn: "Tattoo Artist",      Icon: "🎨"},
	{Key: "food_beverage",    LabelJa: "飲食",                 LabelEn: "Food & Beverage",    Icon: "🍽️"},
	{Key: "bartender",        LabelJa: "バーテンダー",          LabelEn: "Bartender",          Icon: "🍸"},
	{Key: "sommelier",        LabelJa: "ソムリエ",             LabelEn: "Sommelier",          Icon: "🍷"},
	{Key: "personal_trainer", LabelJa: "パーソナルトレーナー",  LabelEn: "Personal Trainer",   Icon: "💪"},
	{Key: "yoga",             LabelJa: "ヨガインストラクター",  LabelEn: "Yoga Instructor",    Icon: "🧘"},
	{Key: "dance",            LabelJa: "ダンスインストラクター", LabelEn: "Dance Instructor",  Icon: "💃"},
	{Key: "photography",      LabelJa: "フォトグラファー",      LabelEn: "Photographer",       Icon: "📷"},
	{Key: "music",            LabelJa: "音楽講師",             LabelEn: "Music Instructor",   Icon: "🎵"},
	{Key: "other",            LabelJa: "その他",               LabelEn: "Other",              Icon: "⭐"},
}

func IsValidJobCategory(key string) bool {
	for _, cat := range JobCategories {
		if cat.Key == key {
			return true
		}
	}
	return false
}
