import React, { useState } from "react";
import t8ns from "../../translations";
import { Lang } from "../../etc";
import Button from "../basic/Button";

type Props = {
  lang: Lang;
  reconnect: () => void;
};

export default ({ lang, reconnect }: Props) => {
  const t8n = t8ns[lang].Disconnected;
  const [disabled, setDisabled] = useState(false);
  return (
    <div className="Disconnected">
      <h1>{t8n.title}</h1>
      <Button
        value={t8n.reconnect}
        onClick={() => {
          setDisabled(true);
          reconnect();
        }}
        disabled={disabled}
      />
    </div>
  );
};
