import React from "react";

import {Menu} from 'antd';
import {NavLink} from "react-router-dom";


export const Navbar = () => {
    const navItems = [
        {
            title: "Туда",
            to: "/"
        },
        {
            title: "Cюда",
            to: "/login"
        }
    ]
    return (
        <Menu mode="horizontal">
            {navItems.map((navItem, index) => {
                const key = index + 1
                return (
                    <Menu.Item key={key}>
                        <NavLink
                            to={navItem.to}
                            exact={index === 0}
                            activeStyle={{
                                color: "red"
                            }}
                        >
                            {navItem.title}
                        </NavLink>
                    </Menu.Item>
                )
            })}
        </Menu>
    )
}