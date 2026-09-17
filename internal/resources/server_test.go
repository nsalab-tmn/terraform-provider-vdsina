// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package resources

import "testing"

func TestServerSchemaValidates(t *testing.T) {
	if err := ServerResource().InternalValidate(nil, true); err != nil {
		t.Fatalf("vdsina_server schema is invalid: %v", err)
	}
}

// Attributes read back from the API must be Computed, or an omitted value forces a replacement on the next plan.
func TestServerReadBackForceNewAttributesAreComputed(t *testing.T) {
	s := ServerResource().Schema
	for _, name := range []string{"template", "ssh_key", "host"} {
		attr, ok := s[name]
		if !ok {
			t.Fatalf("attribute %q missing", name)
		}
		if !attr.Optional || !attr.Computed || !attr.ForceNew {
			t.Errorf("%s: Optional=%v Computed=%v ForceNew=%v, want all true", name, attr.Optional, attr.Computed, attr.ForceNew)
		}
	}
}
