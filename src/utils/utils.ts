export class Utils {
  /**
   * Return a value with a mask applied
   * 
   * ** The return string will be trimed to the same length of the mask
   * @param value value to receive the mask
   * @param mask Every char # will be replaced by the value content
   * @returns 
   * Example of a mask:
   *    
   *  mask: '##.##)##-#$@##'
   * 
   *  result: 12.34)56-7$@8
   */
  static applyMask(value: string, mask: string) {
    const charArray: any = [];
    const maskEntries = mask.split('');
    let splited = value.split('');
    splited = splited.filter(x => mask.indexOf(x) === -1);
    for (const [i, char] of maskEntries.entries()) {

      if (splited.length === 0) {
        break
      }

      if (char === '#') {
        let val = splited.shift();
        if(val) {
          charArray.push(val);
        }
      } else {
        charArray.push(char);
      }
    }

    return charArray.join('');
  }

  // static showShortName(name: string, limit: number) {
  //   if (!limit) {
  //     limit = 10;
  //   }

  //   if (!name && name === '') {
  //     return ''
  //   }

  //   if (name.length >= limit) {
  //     return `${name.slice(0, limit)}...`;
  //   } else {
  //     return name;
  //   }
  // }

  /**
   * Returns true if hex string is a valid objectID from mongodb go driver
   * @param hex Hexadecimal string
   */
  static isValidObjectId(hex: string): boolean {
    try {
      // https://github.com/mongodb/mongo-go-driver/blob/master/bson/primitive/objectid.go
      const n = Number(`0x${hex}`);
      const isString = typeof hex === 'string';
      const is24Long = hex.length === 24;
      const isNumber = !Number.isNaN(n) && n > 0;

      return isString && is24Long && isNumber
    } catch (error) {
      return false;
    }
  }
}