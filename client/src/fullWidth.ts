const fullWidthDigit = (x: string): string => {
  switch (x) {
    case "0":
      return "０";
    case "1":
      return "１";
    case "2":
      return "２";
    case "3":
      return "３";
    case "4":
      return "４";
    case "5":
      return "５";
    case "6":
      return "６";
    case "7":
      return "７";
    case "8":
      return "８";
    case "9":
      return "９";
    default:
      return x;
  }
};

export default (n: number): string =>
  String(n)
    .split("")
    .map(fullWidthDigit)
    .join("");
