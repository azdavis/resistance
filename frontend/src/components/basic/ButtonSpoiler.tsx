import React, { useState, useRef } from "react";
import Button from "./Button";

type Props = {
  view: string;
  spoil: string;
  delay?: number;
};

export default ({ view, spoil, delay = 1000 }: Props) => {
  const [visible, setVisible] = useState(false);
  const id = useRef<NodeJS.Timeout | null>(null);
  return (
    <Button
      value={visible ? spoil : view}
      onClick={() => {
        setVisible(true);
        if (id.current !== null) {
          clearTimeout(id.current);
        }
        id.current = setTimeout(() => setVisible(false), delay);
      }}
    />
  );
};
