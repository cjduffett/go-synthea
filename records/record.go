package records

// Record is a patient's synthesized medical record. This record is generated
// primarilly by modules from the Generic Module Framework.
type Record struct {
	expired       bool
	encounters    []Encounter
	observations  []Observation
	conditions    map[string]Condition
	active        map[string]Condition
	procedures    []Procedure
	immunizations []Immunization
	medications   []Medication
	careplans     []CarePlan
}

/*
func (r *Record) death(time time.Time) {
	r.expired = true
}

func (r *Record) observation(obs Observation, time time.Time) {
	r.observations = append(r.observations, obs)
}

func (r *Record) condition(cond Condition, time time.Time) {
	r.conditions[cond.name]
}

func (r *Record) endCondition(typ string, time time.Time) {
	r.conditions[typ].endTime = time
	r.active[typ] = nil
}

func (r *Record) procedure(typ string, time time.Time, reason string, fhirMethod Procedure, ccdaMethod Procedure) {
	r.active[typ] = Procedure{
		typ:    typ,
		time:   time,
		reason: reason,
		fhir:   fhirMethod,
		ccda:   ccdaMethod,
	}
	r.procedures = append(r.procedures, r.active[typ])
}

func (r *Record) diagnosticReport(typ string, time time.Time, numObs int, fhirMethod DiagnosticReport, ccdaMethod NoAction) {
	r.observations = append(r.observations, Observation{
		typ:    typ,
		time:   time,
		numObs: num_obs,
		fhir:   fhirMethod,
		ccda:   ccdaMethod,
	})
}

// RecordEncounter adds an Encounter to the patient's record.
func (r *Record) RecordEncounter(encounter Encounter, time time.Time) {
	encounter.time = time.Time
	r.encounters = append(r.encounters, encounter)
}

func (r *Record) immunization(typ string, time time.Time, fhirMethod Immunization, ccdaMethod Immunization) {
	r.immunizations = append(r.immunizations, Immunization{
		typ:  imm,
		time: time,
		fhir: fhirMethod,
		ccda: ccdaMethod,
	})
}

func (r *Record) medicationStart(typ string, time time.Time, reasons []string) {
	r.medications = append(r.medications, MedicationOrder{
		typ:        typ,
		time:       time,
		start_time: time,
		reasons:    reasons,
	})
}

func (r *Record) medicationIsActive(med MedicationOrder) bool {
	// See if medication is in this array, and there is no end time
	// r.medications
	return false
}

func (r *Record) updateMedicationReasons(typ string, reasons []string, updateTime time.Time) {
	// Need to find medication typ in the list of medications
	prescription := r.medications[0]
	if prescription {
		// prescription[reasons] = reasons
		// prescription[time] = update_time
	}
}

func (r *Record) medicationStop(typ string, time time.Time, reason string) {
	// Need to find medication typ in the list of medications
	prescription := r.medications[0]
	if prescription {
		// prescription[stop] = time
		// prescription[stop_reason] = reason
	}
}
func (r *Record) careplanStart(typ string, activities []Codes, time time.Time, reason string) {
	r.careplans = append(r.careplans, CarePlanStart{
		typ:        typ,
		activities: activities,
		time:       time,
		start_time: time,
		reasons:    reason,
	})
}
func (r *Record) careplanIsActive(careplan CarePlanStart) bool {
	// Need to find in the list of careplans
	return false
}

func (r *Record) careplanStop(typ string, time time.Time) {
	// Need to find in list of careplans
	careplan := r.careplans[0]
}

func (r *Record) updateCarePlanReasons(typ string, reasons []string, updateTime time.Time) {
	careplan := r.careplans[0]
	if careplan {
		//careplan[reasons] = reasons
		//careplan[time] = update_time
	}
}
*/
