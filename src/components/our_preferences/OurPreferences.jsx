import "./OurPreferences.scss"
import { useState } from "react";
import PreferencesItem from "./preferences_item/PreferencesItem";
import preferences from "./preferences_item/pref_data";

function OurPreferences(){
    const [isVisible, setVisibility] = useState(false);
    return(
    <section className='our-preferences-section'>
            <div className='preferences-section-wrapper'>
                <h2 className="preferences-section-title">Наши преимущества</h2>
                <div className="preferences-list">
                    {preferences.map((preference) => (
                            <PreferencesItem key={preference.id} {...preference} />
                        ))}
                </div>
            </div>

    </section>

    )
}

export default OurPreferences;