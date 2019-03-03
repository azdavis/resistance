import React from "react";
import "./Button.css";

type Props = {
  value: string;
  submit?: boolean;
  disabled?: boolean;
  onClick?: (event: React.MouseEvent<HTMLInputElement, MouseEvent>) => void;
};

export default ({ submit, ...rest }: Props): JSX.Element => (
  <input type={submit ? "submit" : "button"} className="Button" {...rest} />
);
