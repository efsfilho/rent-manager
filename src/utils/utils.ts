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
}