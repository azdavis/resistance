// These should be kept in sync with consts.go.

export const MinN = 5;
export const MaxN = 7;
export const okGameSize = (n: number): boolean => MinN <= n && n <= MaxN;
