// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package regexreplace // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/transformer/regexreplace"

import (
	"context"
	"fmt"
	"regexp"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/entry"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/helper"
)

// Transformer is an operator that will perform a regex-replace on a string.
type Transformer struct {
	helper.TransformerOperator
	field       entry.Field
	regexp      *regexp.Regexp
	replaceWith string
}

// Process will perform a regex-replace on a string
func (t *Transformer) Process(ctx context.Context, entry *entry.Entry) error {
	return t.ProcessWith(ctx, entry, t.replace)
}

// strip will remove ANSI Escape Codes from a string
func (t *Transformer) replace(e *entry.Entry) error {
	value, ok := t.field.Get(e)
	if !ok {
		return nil
	}

	switch v := value.(type) {
	case string:
		s := t.regexp.ReplaceAllString(v, t.replaceWith)
		return t.field.Set(e, s)
	default:
		return fmt.Errorf("type %T cannot be handled", value)
	}
}
