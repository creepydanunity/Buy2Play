import React, { createContext, useState } from 'react';

export const NavbarContext = createContext();

export function NavbarProvider({ children }) {
  const [openIndex, setOpenIndex] = useState(null);

  return (
    <NavbarContext.Provider value={{ openIndex, setOpenIndex }}>
      {children}
    </NavbarContext.Provider>
  );
}
