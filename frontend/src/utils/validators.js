/**
 * Form validation helpers
 * Provides reusable validation functions for form inputs
 */

/**
 * Check if value is empty (null, undefined, empty string, or empty array)
 * @param {*} value - Value to check
 * @returns {boolean}
 */
export function isEmpty(value) {
  if (value === null || value === undefined) return true
  if (typeof value === 'string') return value.trim() === ''
  if (Array.isArray(value)) return value.length === 0
  return false
}

/**
 * Check if value is not empty
 * @param {*} value - Value to check
 * @returns {boolean}
 */
export function isNotEmpty(value) {
  return !isEmpty(value)
}

/**
 * Validate required field
 * @param {*} value - Value to validate
 * @param {string} [fieldName='هذا الحقل'] - Field name for error message
 * @returns {string|null} Error message or null if valid
 */
export function required(value, fieldName = 'هذا الحقل') {
  if (isEmpty(value)) {
    return `${fieldName} مطلوب`
  }
  return null
}

/**
 * Validate minimum length
 * @param {string} value - Value to validate
 * @param {number} min - Minimum length
 * @param {string} [fieldName='هذا الحقل'] - Field name for error message
 * @returns {string|null} Error message or null if valid
 */
export function minLength(value, min, fieldName = 'هذا الحقل') {
  if (isEmpty(value)) return null // Let required handle empty values
  if (value.length < min) {
    return `${fieldName} يجب أن يكون ${min} أحرف على الأقل`
  }
  return null
}

/**
 * Validate maximum length
 * @param {string} value - Value to validate
 * @param {number} max - Maximum length
 * @param {string} [fieldName='هذا الحقل'] - Field name for error message
 * @returns {string|null} Error message or null if valid
 */
export function maxLength(value, max, fieldName = 'هذا الحقل') {
  if (isEmpty(value)) return null
  if (value.length > max) {
    return `${fieldName} يجب أن لا يتجاوز ${max} حرف`
  }
  return null
}

/**
 * Validate email format
 * @param {string} value - Email to validate
 * @returns {string|null} Error message or null if valid
 */
export function email(value) {
  if (isEmpty(value)) return null
  const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/
  if (!emailRegex.test(value)) {
    return 'البريد الإلكتروني غير صالح'
  }
  return null
}

/**
 * Validate Iraqi phone number format
 * @param {string} value - Phone number to validate
 * @returns {string|null} Error message or null if valid
 */
export function phone(value) {
  if (isEmpty(value)) return null
  // Remove spaces and dashes
  const cleaned = value.replace(/[\s-]/g, '')
  // Iraqi phone: 07XX XXX XXXX (11 digits starting with 07)
  const phoneRegex = /^07\d{9}$/
  if (!phoneRegex.test(cleaned)) {
    return 'رقم الهاتف غير صالح (يجب أن يبدأ بـ 07 ويتكون من 11 رقم)'
  }
  return null
}

/**
 * Validate minimum number value
 * @param {number} value - Value to validate
 * @param {number} min - Minimum value
 * @param {string} [fieldName='القيمة'] - Field name for error message
 * @returns {string|null} Error message or null if valid
 */
export function minValue(value, min, fieldName = 'القيمة') {
  if (isEmpty(value)) return null
  if (Number(value) < min) {
    return `${fieldName} يجب أن تكون ${min} على الأقل`
  }
  return null
}

/**
 * Validate maximum number value
 * @param {number} value - Value to validate
 * @param {number} max - Maximum value
 * @param {string} [fieldName='القيمة'] - Field name for error message
 * @returns {string|null} Error message or null if valid
 */
export function maxValue(value, max, fieldName = 'القيمة') {
  if (isEmpty(value)) return null
  if (Number(value) > max) {
    return `${fieldName} يجب أن لا تتجاوز ${max}`
  }
  return null
}

/**
 * Validate that value is a positive number
 * @param {number} value - Value to validate
 * @param {string} [fieldName='القيمة'] - Field name for error message
 * @returns {string|null} Error message or null if valid
 */
export function positiveNumber(value, fieldName = 'القيمة') {
  if (isEmpty(value)) return null
  if (Number(value) <= 0) {
    return `${fieldName} يجب أن تكون رقم موجب`
  }
  return null
}

/**
 * Validate that value is a non-negative number
 * @param {number} value - Value to validate
 * @param {string} [fieldName='القيمة'] - Field name for error message
 * @returns {string|null} Error message or null if valid
 */
export function nonNegative(value, fieldName = 'القيمة') {
  if (isEmpty(value)) return null
  if (Number(value) < 0) {
    return `${fieldName} يجب أن لا تكون سالبة`
  }
  return null
}

/**
 * Validate date format
 * @param {string} value - Date string to validate
 * @returns {string|null} Error message or null if valid
 */
export function date(value) {
  if (isEmpty(value)) return null
  const dateObj = new Date(value)
  if (isNaN(dateObj.getTime())) {
    return 'التاريخ غير صالح'
  }
  return null
}

/**
 * Validate that date is not in the past
 * @param {string} value - Date string to validate
 * @returns {string|null} Error message or null if valid
 */
export function futureDate(value) {
  if (isEmpty(value)) return null
  const dateObj = new Date(value)
  const today = new Date()
  today.setHours(0, 0, 0, 0)

  if (dateObj < today) {
    return 'التاريخ يجب أن يكون في المستقبل'
  }
  return null
}

/**
 * Validate that date is not in the future
 * @param {string} value - Date string to validate
 * @returns {string|null} Error message or null if valid
 */
export function pastDate(value) {
  if (isEmpty(value)) return null
  const dateObj = new Date(value)
  const today = new Date()
  today.setHours(23, 59, 59, 999)

  if (dateObj > today) {
    return 'التاريخ يجب أن يكون في الماضي'
  }
  return null
}

/**
 * Validate password strength
 * @param {string} value - Password to validate
 * @returns {string|null} Error message or null if valid
 */
export function password(value) {
  if (isEmpty(value)) return null
  if (value.length < 8) {
    return 'كلمة المرور يجب أن تكون 8 أحرف على الأقل'
  }
  return null
}

/**
 * Validate password confirmation matches
 * @param {string} value - Password confirmation
 * @param {string} originalPassword - Original password to match
 * @returns {string|null} Error message or null if valid
 */
export function passwordMatch(value, originalPassword) {
  if (isEmpty(value)) return null
  if (value !== originalPassword) {
    return 'كلمة المرور غير متطابقة'
  }
  return null
}

/**
 * Validate username format (alphanumeric, underscores)
 * @param {string} value - Username to validate
 * @returns {string|null} Error message or null if valid
 */
export function username(value) {
  if (isEmpty(value)) return null
  const usernameRegex = /^[a-zA-Z0-9_]+$/
  if (!usernameRegex.test(value)) {
    return 'اسم المستخدم يجب أن يحتوي على أحرف وأرقام وشرطات سفلية فقط'
  }
  if (value.length < 3) {
    return 'اسم المستخدم يجب أن يكون 3 أحرف على الأقل'
  }
  return null
}

/**
 * Validate percentage (0-100)
 * @param {number} value - Value to validate
 * @returns {string|null} Error message or null if valid
 */
export function percentage(value) {
  if (isEmpty(value)) return null
  const num = Number(value)
  if (num < 0 || num > 100) {
    return 'النسبة يجب أن تكون بين 0 و 100'
  }
  return null
}

/**
 * Run multiple validators on a value
 * @param {*} value - Value to validate
 * @param {Function[]} validators - Array of validator functions
 * @returns {string|null} First error message or null if all valid
 */
export function validate(value, validators) {
  for (const validator of validators) {
    const error = validator(value)
    if (error) return error
  }
  return null
}

/**
 * Validate an entire form object
 * @param {Object} form - Form data object
 * @param {Object} rules - Validation rules (field -> array of validators)
 * @returns {Object} Object with field names as keys and error messages as values
 */
export function validateForm(form, rules) {
  const errors = {}

  for (const [field, validators] of Object.entries(rules)) {
    const error = validate(form[field], validators)
    if (error) {
      errors[field] = error
    }
  }

  return errors
}

/**
 * Check if form validation has any errors
 * @param {Object} errors - Errors object from validateForm
 * @returns {boolean}
 */
export function hasErrors(errors) {
  return Object.keys(errors).length > 0
}

export default {
  isEmpty,
  isNotEmpty,
  required,
  minLength,
  maxLength,
  email,
  phone,
  minValue,
  maxValue,
  positiveNumber,
  nonNegative,
  date,
  futureDate,
  pastDate,
  password,
  passwordMatch,
  username,
  percentage,
  validate,
  validateForm,
  hasErrors,
}
