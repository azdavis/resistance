import React from "react";
import "./Button.css";
import "./FullWidth.css";

type Props = {
  value: string;
  type?: "button" | "submit" | "reset";
  disabled?: boolean;
  onClick?: (event: React.MouseEvent<HTMLInputElement, MouseEvent>) => void;
};

export default ({ type = "button", ...rest }: Props) => (
  <input type={type} className="Button FullWidth" {...rest} />
);
