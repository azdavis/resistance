// These should be kept in sync with consts.go.

export const minN = 5;
export const maxN = 7;
export const okGameSize = (n: number): boolean => minN <= n && n <= maxN;
export const maxPts = 3;
export const maxSkip = 3;
