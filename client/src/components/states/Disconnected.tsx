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
      {t8n.title}
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
