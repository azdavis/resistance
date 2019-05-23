import { Lang, langs } from "./etc";

const langKey = "lang";

export default {
  getLang: (): Lang | null => {
    const x = localStorage.getItem(langKey);
    return x !== null && (langs as Array<string>).includes(x)
      ? (x as Lang)
      : null;
  },
  setLang: (x: Lang) => localStorage.setItem(langKey, x),
};
