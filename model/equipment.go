package model

type Stat struct {
	Type struct {
		Type string `json:"type"`
		Name string `json:"name"`
	} `json:"type"`
	Value   int `json:"value"`
	Display struct {
		DisplayString string `json:"display_string"`
		Color         struct {
			R int `json:"r"`
			G int `json:"g"`
			B int `json:"b"`
			A int `json:"a"`
		} `json:"color"`
	} `json:"display"`
	IsNegated    bool `json:"is_negated,omitempty"`
	IsEquipBonus bool `json:"is_equip_bonus,omitempty"`
}
type PlayableClassesLink struct {
	Key struct {
		Href string `json:"href"`
	} `json:"key"`
	Name string `json:"name"`
	Id   int    `json:"id"`
}
type Item struct {
	Item struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Name string `json:"name"`
		Id   int    `json:"id"`
	} `json:"item"`
	IsEquipped bool `json:"is_equipped"`
}
type SetEffect struct {
	DisplayString string `json:"display_string"`
	RequiredCount int    `json:"required_count"`
	IsActive      bool   `json:"is_active"`
}
type ModifiedCraftingStat struct {
	Id   int    `json:"id"`
	Type string `json:"type"`
	Name string `json:"name"`
}
type Enchantment struct {
	DisplayString   string `json:"display_string"`
	EnchantmentId   int    `json:"enchantment_id"`
	EnchantmentSlot struct {
		Id   int    `json:"id"`
		Type string `json:"type"`
	} `json:"enchantment_slot"`
	SourceItem struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Name string `json:"name"`
		Id   int    `json:"id"`
	} `json:"source_item,omitempty"`
	Spell struct {
		Spell struct {
			Key struct {
				Href string `json:"href"`
			} `json:"key"`
			Name string `json:"name"`
			Id   int    `json:"id"`
		} `json:"spell"`
		Description string `json:"description"`
	} `json:"spell,omitempty"`
}

type Spell struct {
	Spell struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Name string `json:"name"`
		Id   int    `json:"id"`
	} `json:"spell"`
	Description  string `json:"description"`
	DisplayColor struct {
		R int `json:"r"`
		G int `json:"g"`
		B int `json:"b"`
		A int `json:"a"`
	} `json:"display_color,omitempty"`
}
type EquippedItem struct {
	Item struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Id int `json:"id"`
	} `json:"item"`
	Sockets []struct {
		SocketType struct {
			Type string `json:"type"`
			Name string `json:"name"`
		} `json:"socket_type"`
		Item struct {
			Key struct {
				Href string `json:"href"`
			} `json:"key"`
			Name string `json:"name"`
			Id   int    `json:"id"`
		} `json:"item"`
		DisplayString string `json:"display_string"`
		Media         struct {
			Key struct {
				Href string `json:"href"`
			} `json:"key"`
			Id int `json:"id"`
		} `json:"media"`
		Context int `json:"context,omitempty"`
	} `json:"sockets,omitempty"`
	Slot struct {
		Type string `json:"type"`
		Name string `json:"name"`
	} `json:"slot"`
	Quantity  int   `json:"quantity"`
	Context   int   `json:"context,omitempty"`
	BonusList []int `json:"bonus_list,omitempty"`
	Quality   struct {
		Type string `json:"type"`
		Name string `json:"name"`
	} `json:"quality"`
	Name                 string `json:"name"`
	ModifiedAppearanceId int    `json:"modified_appearance_id,omitempty"`
	Media                struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Id int `json:"id"`
	} `json:"media"`
	ItemClass struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Name string `json:"name"`
		Id   int    `json:"id"`
	} `json:"item_class"`
	ItemSubclass struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Name string `json:"name"`
		Id   int    `json:"id"`
	} `json:"item_subclass"`
	InventoryType struct {
		Type string `json:"type"`
		Name string `json:"name"`
	} `json:"inventory_type"`
	Binding struct {
		Type string `json:"type"`
		Name string `json:"name"`
	} `json:"binding"`
	Armor struct {
		Value   int `json:"value"`
		Display struct {
			DisplayString string `json:"display_string"`
			Color         struct {
				R int `json:"r"`
				G int `json:"g"`
				B int `json:"b"`
				A int `json:"a"`
			} `json:"color"`
		} `json:"display"`
	} `json:"armor,omitempty"`
	Stats     []Stat `json:"stats,omitempty"`
	SellPrice struct {
		Value          int `json:"value"`
		DisplayStrings struct {
			Header string `json:"header"`
			Gold   string `json:"gold"`
			Silver string `json:"silver"`
			Copper string `json:"copper"`
		} `json:"display_strings"`
	} `json:"sell_price,omitempty"`
	Requirements struct {
		Level struct {
			Value         int    `json:"value"`
			DisplayString string `json:"display_string"`
		} `json:"level"`
		PlayableClasses struct {
			Links         []PlayableClassesLink `json:"links"`
			DisplayString string                `json:"display_string"`
		} `json:"playable_classes,omitempty"`
	} `json:"requirements,omitempty"`
	Set struct {
		ItemSet struct {
			Key struct {
				Href string `json:"href"`
			} `json:"key"`
			Name string `json:"name"`
			Id   int    `json:"id"`
		} `json:"item_set"`
		Items         []Item      `json:"items"`
		Effects       []SetEffect `json:"effects"`
		DisplayString string      `json:"display_string"`
	} `json:"set,omitempty"`
	Level struct {
		Value         int    `json:"value"`
		DisplayString string `json:"display_string"`
	} `json:"level"`
	Transmog struct {
		Item struct {
			Key struct {
				Href string `json:"href"`
			} `json:"key"`
			Name string `json:"name"`
			Id   int    `json:"id"`
		} `json:"item"`
		DisplayString            string `json:"display_string"`
		ItemModifiedAppearanceId int    `json:"item_modified_appearance_id"`
	} `json:"transmog,omitempty"`
	Durability struct {
		Value         int    `json:"value"`
		DisplayString string `json:"display_string"`
	} `json:"durability,omitempty"`
	NameDescription struct {
		DisplayString string `json:"display_string"`
		Color         struct {
			R int `json:"r"`
			G int `json:"g"`
			B int `json:"b"`
			A int `json:"a"`
		} `json:"color"`
	} `json:"name_description,omitempty"`
	IsSubclassHidden     bool                   `json:"is_subclass_hidden,omitempty"`
	ModifiedCraftingStat []ModifiedCraftingStat `json:"modified_crafting_stat,omitempty"`
	Enchantments         []Enchantment          `json:"enchantments,omitempty"`
	LimitCategory        string                 `json:"limit_category,omitempty"`
	Spells               []Spell                `json:"spells,omitempty"`
	UniqueEquipped       string                 `json:"unique_equipped,omitempty"`
	Weapon               struct {
		Damage struct {
			MinValue      int    `json:"min_value"`
			MaxValue      int    `json:"max_value"`
			DisplayString string `json:"display_string"`
			DamageClass   struct {
				Type string `json:"type"`
				Name string `json:"name"`
			} `json:"damage_class"`
		} `json:"damage"`
		AttackSpeed struct {
			Value         int    `json:"value"`
			DisplayString string `json:"display_string"`
		} `json:"attack_speed"`
		Dps struct {
			Value         float64 `json:"value"`
			DisplayString string  `json:"display_string"`
		} `json:"dps"`
	} `json:"weapon,omitempty"`
}

type EquippedItemSet struct {
	ItemSet struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Name string `json:"name"`
		Id   int    `json:"id"`
	} `json:"item_set"`
	Items         []Item      `json:"items"`
	Effects       []SetEffect `json:"effects"`
	DisplayString string      `json:"display_string"`
}
type Equipment struct {
	Links     Links `json:"_links"`
	Character struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Name  string `json:"name"`
		Id    int    `json:"id"`
		Realm struct {
			Key struct {
				Href string `json:"href"`
			} `json:"key"`
			Name string `json:"name"`
			Id   int    `json:"id"`
			Slug string `json:"slug"`
		} `json:"realm"`
	} `json:"character"`
	EquippedItems    []EquippedItem    `json:"equipped_items"`
	EquippedItemSets []EquippedItemSet `json:"equipped_item_sets"`
}
