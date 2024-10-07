package utils

//func Wrap(err error) error {
//
//	switch {
//	case err == nil:
//		return nil
//	case errors.Is(err, gorm.ErrRecordNotFound):
//		return errors.Wrap(domain.ErrNotFound, err.Error())
//	case errors.Is(err, gorm.ErrDuplicatedKey):
//		return errors.Wrap(domain.ErrUniqueKeyViolation, err.Error())
//	case errors.Is(err, gorm.ErrForeignKeyViolated):
//		return errors.Wrap(domain.ErrForeignKeyViolation, err.Error())
//	case errors.Is(err, gorm.ErrCheckConstraintViolated):
//		return errors.Wrap(domain.ErrCheckConstraint, err.Error())
//	case errors.Is(err, gorm.ErrInvalidTransaction):
//		return errors.Wrap(domain.ErrInvalidTransaction, err.Error())
//	case errors.Is(err, gorm.ErrInvalidData),
//		errors.Is(err, gorm.ErrInvalidField),
//		errors.Is(err, gorm.ErrInvalidValue),
//		errors.Is(err, gorm.ErrInvalidValueOfLength):
//		return errors.Wrap(domain.ErrBadParam, err.Error())
//	case errors.Is(err, gorm.ErrMissingWhereClause),
//		errors.Is(err, gorm.ErrPrimaryKeyRequired),
//		errors.Is(err, gorm.ErrModelValueRequired),
//		errors.Is(err, gorm.ErrModelAccessibleFieldsRequired),
//		errors.Is(err, gorm.ErrSubQueryRequired):
//		return errors.Wrap(domain.ErrInvalidQuery, err.Error())
//	case errors.Is(err, gorm.ErrNotImplemented),
//		errors.Is(err, gorm.ErrUnsupportedRelation),
//		errors.Is(err, gorm.ErrUnsupportedDriver),
//		errors.Is(err, gorm.ErrDryRunModeUnsupported):
//		return errors.Wrap(domain.ErrUnsupportedOperation, err.Error())
//	case errors.Is(err, gorm.ErrRegistered),
//		errors.Is(err, gorm.ErrEmptySlice),
//		errors.Is(err, gorm.ErrInvalidDB),
//		errors.Is(err, gorm.ErrPreloadNotAllowed):
//		return errors.Wrap(domain.ErrDatabaseOperation, err.Error())
//	default:
//		return errors.Wrap(domain.ErrInternalServerError, err.Error())
//	}
//}
