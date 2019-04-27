import { Lang, langs } from "./shared";

const langKey = "lang";

export const getLang = (): Lang | null => {
  const x = localStorage.getItem(langKey);
  return x !== null && (langs as Array<string>).includes(x)
    ? (x as Lang)
    : null;
};

export const setLang = (x: Lang) => localStorage.setItem(langKey, x);
