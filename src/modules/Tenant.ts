
import moment from "moment";

// export interface Tenant {
//   name: string;
//   cpf: string;
//   rg: string;
//   birthDate: number;
// }

interface Validator {

  /**
   * Checks the name has minimum 4 and maximum 50 letters(disconsidering spaces)
   */
  isNameValid(name: string): boolean;

  /**
   * Checks if it's a valid cpf number
   */
  isCpfValid(cpf: string): boolean;

  /**
   * Checks if the number rg digits is between 8 and 12 digits.
   */
  isRgValid(rg: string): boolean;

  /**
   * Checks if it's a valid date with DD/MM/YYYY format.
   * (Future dates not allowed)
   */
  isBirthDateValid(date: string): boolean;
}

export class TenantValidator implements Validator {

  isNameValid(name: string) {
    const nameLength = name.replaceAll(/\s/g ,'').length;
    return nameLength >= 4 && nameLength <= 50;
  }

  isCpfValid(cpf: string) {
   
    cpf = cpf.replaceAll(/[^0-9]/g,'');

    if(cpf == '' || cpf.length != 11) {
      return false;
    }
    
    let knowInvalid = [cpf].some((n) => {
      for (let i = 0; i < 10; i++) {
        if(n === `${i}`.repeat(11)) {
          return true;
        }
      }
    });

    // Elimina CPFs inválidos conhecidos
    if(knowInvalid) {
      return false
    }

    // Valida 1º digito	
    let add = 0;
    for (let i=0; i < 9; i ++) {
      add += parseInt(cpf[i]) * (10 - i);
    }
    let rev = 11 - (add % 11);
    if (rev == 10 || rev == 11) {
      rev = 0;
    }
    if (rev != parseInt(cpf[9])) {
      return false;
    }

    // Valida 2º digito
    add = 0;
    for (let i = 0; i < 10; i ++) {
      add += parseInt(cpf[i]) * (11 - i);
    }
    rev = 11 - (add % 11);
    if (rev == 10 || rev == 11) {
      rev = 0;
    }
    if (rev != parseInt(cpf[10])) {
      return false;
    }

    return true;
  }

  isRgValid(rg: string) {
    const digits = rg.replaceAll(/[^\d]/g,'');
    return digits.length >= 8 && digits.length <= 12;
  }

  isBirthDateValid(date: string) {
    let momentDate = moment(date, "DD/MM/YYYY", true);
    return momentDate.isValid() && momentDate.isBefore(moment());
  }
}