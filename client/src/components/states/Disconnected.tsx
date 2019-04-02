import React, { useState } from "react";
import { Lang } from "../../types";
import Button from "../basic/Button";

type Props = {
  lang: Lang;
  reconnect: () => void;
};

const text = {
  title: {
    en: <h1>Disconnected</h1>,
  },
  reconnect: {
    en: "Reconnect",
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
