export const writeSessionStorage = (key: string, value: string) => {
  try {
    const jsonValue = JSON.stringify(value)
    sessionStorage.setItem(key, jsonValue)
  } catch (e) {
    console.error('Error writing session storage', key, value)
  }
}

export const parseStorageItem = (storage: any, key: string, defaultValue?: any) => {
  try {
    const value = storage.getItem(key)
    if (!value) {
      return defaultValue
    }
    const jsonValue = JSON.parse(value)
    return jsonValue
  } catch (e) {
    console.error('error parsing JSON item', e)
    return defaultValue
  }
}


