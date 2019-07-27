import React from "react";
import "./Checkbox.css";

type Props = {
  checked?: boolean;
  disabled?: boolean;
  onChange?: (e: React.ChangeEvent<HTMLInputElement>) => void;
};

export default (props: Props) => (
  <input className="Checkbox" type="checkbox" {...props} />
);
