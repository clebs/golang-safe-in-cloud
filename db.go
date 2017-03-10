package sic

// SetSaveCallback sets a callback which will be triggered
// when the database was changed over the functions
func (c *Database) SetSaveCallback(fnc func()) {
	c.Callback = fnc
}

// AddCard adds a card to the model
func (c *Database) AddCard(card Card) {
	c.Cards = append(c.Cards, card)
	c.checkCallback()
}

// AddLabel adds an label to the model
func (c *Database) AddLabel(label Label) {
	c.Labels = append(c.Labels, label)
	c.checkCallback()
}

// checkCallback checks if the callback was set and
// triggers them if
func (c *Database) checkCallback() {
	if c.Callback != nil {
		c.Callback()
	}
}
