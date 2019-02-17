import { Dispatch } from "react";

type Msg = { t: "nameChoose"; name: string };

export type Send = Dispatch<Msg>;
