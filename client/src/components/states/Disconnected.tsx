import React, { useState } from "react";
import { Lang } from "../../etc";
import Button from "../basic/Button";

type Props = {
  lang: Lang;
  reconnect: () => void;
};

const text = {
  title: {
    en: <h1>Disconnected</h1>,
    ja: <h1>接続が切られた</h1>,
  },
  reconnect: {
    en: "Reconnect",
    ja: "再接続する",
  },
};

export default ({ lang, reconnect }: Props) => {
  const [disabled, setDisabled] = useState(false);
  return (
    <div className="Disconnected">
      {text.title[lang]}
      <Button
        value={text.reconnect[lang]}
        onClick={() => {
          setDisabled(true);
          reconnect();
        }}
        disabled={disabled}
      />
    </div>
  );
};
