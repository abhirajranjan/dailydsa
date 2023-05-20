package permissions

// permissions that is exported
type Permissions interface {
	Has(Permissions) bool
}

// used for setting single permission
type permission int

const (
	// user roles
	User permission = iota
	Admin

	// system roles
	Frontend
)

func (p permission) Has(perm Permissions) bool {
	if tmultiperm, e := perm.(multipermission); e {
		// perm is of type multipermision
		_, exists := tmultiperm[p]
		return exists
	}

	if tperm, e := perm.(permission); e {
		// perm if of type permission
		return p == tperm
	}

	return false
}

// used for setting multiple permissions
type multipermission map[permission]int

func (m multipermission) Has(perm Permissions) bool {
	if tmultiperm, e := perm.(multipermission); e {
		for key := range tmultiperm {
			if _, exists := m[key]; exists {
				return true
			}
		}
		return false
	}

	if tperm, e := perm.(permission); e {
		_, exists := m[tperm]
		return exists
	}

	return false
}

func MultiPermission(perms ...permission) (multiperm multipermission) {
	multiperm = make(multipermission)
	for i, perm := range perms {
		multiperm[perm] = i
	}
	return
}
