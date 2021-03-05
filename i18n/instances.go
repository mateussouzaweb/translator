package i18n

var _instances []*Context

// LoadContext method
func LoadContext(context *Context) error {

	err := context.Load()

	if err != nil {
		return err
	}

	_instances = append(_instances, context)

	return nil
}

// GetContext return context language instance the create translations
func GetContext(code string, fallback string) *Context {

	context := &Context{}

	// Match context
	if code != "" {
		for _, item := range _instances {
			if item.Code == code || item.Alias == code {
				context = item
			}
		}
	}

	// Fallback context
	if context.Code == "" && fallback != "" {
		for _, item := range _instances {
			if item.Code == fallback || item.Alias == fallback {
				context = item
			}
		}
	}

	return context
}
