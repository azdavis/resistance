import React from "react";
import "./TextInput.css";

export default React.forwardRef<HTMLInputElement>((_, ref) => (
  <input className="TextInput" type="text" autoCorrect="off" ref={ref} />
));
