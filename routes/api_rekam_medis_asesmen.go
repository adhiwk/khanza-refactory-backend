// Kode dibangkitkan dari sik.sql + source/src/rekammedis; satu baris per form asesmen.
package routes

import (
	"github.com/goravel/framework/contracts/route"

	rmaction "goravel/app/actions/rekammedis"
	rmctrl "goravel/app/http/controllers/rekammedis"
	rmreq "goravel/app/http/requests/rekammedis"
	"goravel/app/modules/rbac"
)

// registerRekamMedisAsesmenRoutes /rekam-medis/<form>: GET (list), GET /detail, GET /cetak, POST, PUT, DELETE; kunci lewat query string.
func registerRekamMedisAsesmenRoutes(router route.Router) {
	perm := rbac.RequirePermission
	reg := func(slug string, c rmctrl.Handler) {
		p := "/rekam-medis/" + slug
		router.Middleware(perm("rekam_medis.view")).Get(p, c.Index)
		router.Middleware(perm("rekam_medis.view")).Get(p+"/detail", c.Show)
		router.Middleware(perm("rekam_medis.create")).Post(p, c.Store)
		router.Middleware(perm("rekam_medis.update")).Put(p, c.Update)
		router.Middleware(perm("rekam_medis.delete")).Delete(p, c.Destroy)
		router.Middleware(perm("rekam_medis.view")).Get(p+"/cetak", c.Cetak)
	}
	reg("admisi-skoring-tolac", rmctrl.NewController(rmaction.FormAdmisiSkoringTolac, func() rmreq.StoreRequest[rmreq.AdmisiSkoringTolacData] { return &rmreq.AdmisiSkoringTolacStore{} }, func() rmreq.UpdateRequest[rmreq.AdmisiSkoringTolacData] { return &rmreq.AdmisiSkoringTolacUpdate{} }))
	reg("catatan-adime-gizi", rmctrl.NewController(rmaction.FormCatatanAdimeGizi, func() rmreq.StoreRequest[rmreq.CatatanAdimeGiziData] { return &rmreq.CatatanAdimeGiziStore{} }, func() rmreq.UpdateRequest[rmreq.CatatanAdimeGiziData] { return &rmreq.CatatanAdimeGiziUpdate{} }))
	reg("catatan-anestesi-sedasi", rmctrl.NewController(rmaction.FormCatatanAnestesiSedasi, func() rmreq.StoreRequest[rmreq.CatatanAnestesiSedasiData] { return &rmreq.CatatanAnestesiSedasiStore{} }, func() rmreq.UpdateRequest[rmreq.CatatanAnestesiSedasiData] {
		return &rmreq.CatatanAnestesiSedasiUpdate{}
	}))
	reg("catatan-pengkajian-paska-operasi", rmctrl.NewController(rmaction.FormCatatanPengkajianPaskaOperasi, func() rmreq.StoreRequest[rmreq.CatatanPengkajianPaskaOperasiData] {
		return &rmreq.CatatanPengkajianPaskaOperasiStore{}
	}, func() rmreq.UpdateRequest[rmreq.CatatanPengkajianPaskaOperasiData] {
		return &rmreq.CatatanPengkajianPaskaOperasiUpdate{}
	}))
	reg("catatan-persalinan", rmctrl.NewController(rmaction.FormCatatanPersalinan, func() rmreq.StoreRequest[rmreq.CatatanPersalinanData] { return &rmreq.CatatanPersalinanStore{} }, func() rmreq.UpdateRequest[rmreq.CatatanPersalinanData] { return &rmreq.CatatanPersalinanUpdate{} }))
	reg("checklist-kesiapan-anestesi", rmctrl.NewController(rmaction.FormChecklistKesiapanAnestesi, func() rmreq.StoreRequest[rmreq.ChecklistKesiapanAnestesiData] {
		return &rmreq.ChecklistKesiapanAnestesiStore{}
	}, func() rmreq.UpdateRequest[rmreq.ChecklistKesiapanAnestesiData] {
		return &rmreq.ChecklistKesiapanAnestesiUpdate{}
	}))
	reg("checklist-kriteria-keluar-hcu", rmctrl.NewController(rmaction.FormChecklistKriteriaKeluarHcu, func() rmreq.StoreRequest[rmreq.ChecklistKriteriaKeluarHcuData] {
		return &rmreq.ChecklistKriteriaKeluarHcuStore{}
	}, func() rmreq.UpdateRequest[rmreq.ChecklistKriteriaKeluarHcuData] {
		return &rmreq.ChecklistKriteriaKeluarHcuUpdate{}
	}))
	reg("checklist-kriteria-keluar-icu", rmctrl.NewController(rmaction.FormChecklistKriteriaKeluarIcu, func() rmreq.StoreRequest[rmreq.ChecklistKriteriaKeluarIcuData] {
		return &rmreq.ChecklistKriteriaKeluarIcuStore{}
	}, func() rmreq.UpdateRequest[rmreq.ChecklistKriteriaKeluarIcuData] {
		return &rmreq.ChecklistKriteriaKeluarIcuUpdate{}
	}))
	reg("checklist-kriteria-keluar-isolasi", rmctrl.NewController(rmaction.FormChecklistKriteriaKeluarIsolasi, func() rmreq.StoreRequest[rmreq.ChecklistKriteriaKeluarIsolasiData] {
		return &rmreq.ChecklistKriteriaKeluarIsolasiStore{}
	}, func() rmreq.UpdateRequest[rmreq.ChecklistKriteriaKeluarIsolasiData] {
		return &rmreq.ChecklistKriteriaKeluarIsolasiUpdate{}
	}))
	reg("checklist-kriteria-keluar-nicu", rmctrl.NewController(rmaction.FormChecklistKriteriaKeluarNicu, func() rmreq.StoreRequest[rmreq.ChecklistKriteriaKeluarNicuData] {
		return &rmreq.ChecklistKriteriaKeluarNicuStore{}
	}, func() rmreq.UpdateRequest[rmreq.ChecklistKriteriaKeluarNicuData] {
		return &rmreq.ChecklistKriteriaKeluarNicuUpdate{}
	}))
	reg("checklist-kriteria-keluar-picu", rmctrl.NewController(rmaction.FormChecklistKriteriaKeluarPicu, func() rmreq.StoreRequest[rmreq.ChecklistKriteriaKeluarPicuData] {
		return &rmreq.ChecklistKriteriaKeluarPicuStore{}
	}, func() rmreq.UpdateRequest[rmreq.ChecklistKriteriaKeluarPicuData] {
		return &rmreq.ChecklistKriteriaKeluarPicuUpdate{}
	}))
	reg("checklist-kriteria-masuk-hcu", rmctrl.NewController(rmaction.FormChecklistKriteriaMasukHcu, func() rmreq.StoreRequest[rmreq.ChecklistKriteriaMasukHcuData] {
		return &rmreq.ChecklistKriteriaMasukHcuStore{}
	}, func() rmreq.UpdateRequest[rmreq.ChecklistKriteriaMasukHcuData] {
		return &rmreq.ChecklistKriteriaMasukHcuUpdate{}
	}))
	reg("checklist-kriteria-masuk-icu", rmctrl.NewController(rmaction.FormChecklistKriteriaMasukIcu, func() rmreq.StoreRequest[rmreq.ChecklistKriteriaMasukIcuData] {
		return &rmreq.ChecklistKriteriaMasukIcuStore{}
	}, func() rmreq.UpdateRequest[rmreq.ChecklistKriteriaMasukIcuData] {
		return &rmreq.ChecklistKriteriaMasukIcuUpdate{}
	}))
	reg("checklist-kriteria-masuk-isolasi", rmctrl.NewController(rmaction.FormChecklistKriteriaMasukIsolasi, func() rmreq.StoreRequest[rmreq.ChecklistKriteriaMasukIsolasiData] {
		return &rmreq.ChecklistKriteriaMasukIsolasiStore{}
	}, func() rmreq.UpdateRequest[rmreq.ChecklistKriteriaMasukIsolasiData] {
		return &rmreq.ChecklistKriteriaMasukIsolasiUpdate{}
	}))
	reg("checklist-kriteria-masuk-nicu", rmctrl.NewController(rmaction.FormChecklistKriteriaMasukNicu, func() rmreq.StoreRequest[rmreq.ChecklistKriteriaMasukNicuData] {
		return &rmreq.ChecklistKriteriaMasukNicuStore{}
	}, func() rmreq.UpdateRequest[rmreq.ChecklistKriteriaMasukNicuData] {
		return &rmreq.ChecklistKriteriaMasukNicuUpdate{}
	}))
	reg("checklist-kriteria-masuk-picu", rmctrl.NewController(rmaction.FormChecklistKriteriaMasukPicu, func() rmreq.StoreRequest[rmreq.ChecklistKriteriaMasukPicuData] {
		return &rmreq.ChecklistKriteriaMasukPicuStore{}
	}, func() rmreq.UpdateRequest[rmreq.ChecklistKriteriaMasukPicuData] {
		return &rmreq.ChecklistKriteriaMasukPicuUpdate{}
	}))
	reg("checklist-pemberian-fibrinolitik", rmctrl.NewController(rmaction.FormChecklistPemberianFibrinolitik, func() rmreq.StoreRequest[rmreq.ChecklistPemberianFibrinolitikData] {
		return &rmreq.ChecklistPemberianFibrinolitikStore{}
	}, func() rmreq.UpdateRequest[rmreq.ChecklistPemberianFibrinolitikData] {
		return &rmreq.ChecklistPemberianFibrinolitikUpdate{}
	}))
	reg("checklist-post-operasi", rmctrl.NewController(rmaction.FormChecklistPostOperasi, func() rmreq.StoreRequest[rmreq.ChecklistPostOperasiData] { return &rmreq.ChecklistPostOperasiStore{} }, func() rmreq.UpdateRequest[rmreq.ChecklistPostOperasiData] { return &rmreq.ChecklistPostOperasiUpdate{} }))
	reg("checklist-pre-operasi", rmctrl.NewController(rmaction.FormChecklistPreOperasi, func() rmreq.StoreRequest[rmreq.ChecklistPreOperasiData] { return &rmreq.ChecklistPreOperasiStore{} }, func() rmreq.UpdateRequest[rmreq.ChecklistPreOperasiData] { return &rmreq.ChecklistPreOperasiUpdate{} }))
	reg("asuhan-gizi", rmctrl.NewController(rmaction.FormAsuhanGizi, func() rmreq.StoreRequest[rmreq.AsuhanGiziData] { return &rmreq.AsuhanGiziStore{} }, func() rmreq.UpdateRequest[rmreq.AsuhanGiziData] { return &rmreq.AsuhanGiziUpdate{} }))
	reg("catatan-cairan-hemodialisa", rmctrl.NewController(rmaction.FormCatatanCairanHemodialisa, func() rmreq.StoreRequest[rmreq.CatatanCairanHemodialisaData] {
		return &rmreq.CatatanCairanHemodialisaStore{}
	}, func() rmreq.UpdateRequest[rmreq.CatatanCairanHemodialisaData] {
		return &rmreq.CatatanCairanHemodialisaUpdate{}
	}))
	reg("catatan-cek-gds", rmctrl.NewController(rmaction.FormCatatanCekGds, func() rmreq.StoreRequest[rmreq.CatatanCekGdsData] { return &rmreq.CatatanCekGdsStore{} }, func() rmreq.UpdateRequest[rmreq.CatatanCekGdsData] { return &rmreq.CatatanCekGdsUpdate{} }))
	reg("catatan-keperawatan-ralan", rmctrl.NewController(rmaction.FormCatatanKeperawatanRalan, func() rmreq.StoreRequest[rmreq.CatatanKeperawatanRalanData] {
		return &rmreq.CatatanKeperawatanRalanStore{}
	}, func() rmreq.UpdateRequest[rmreq.CatatanKeperawatanRalanData] {
		return &rmreq.CatatanKeperawatanRalanUpdate{}
	}))
	reg("catatan-keperawatan-ranap", rmctrl.NewController(rmaction.FormCatatanKeperawatanRanap, func() rmreq.StoreRequest[rmreq.CatatanKeperawatanRanapData] {
		return &rmreq.CatatanKeperawatanRanapStore{}
	}, func() rmreq.UpdateRequest[rmreq.CatatanKeperawatanRanapData] {
		return &rmreq.CatatanKeperawatanRanapUpdate{}
	}))
	reg("catatan-keseimbangan-cairan", rmctrl.NewController(rmaction.FormCatatanKeseimbanganCairan, func() rmreq.StoreRequest[rmreq.CatatanKeseimbanganCairanData] {
		return &rmreq.CatatanKeseimbanganCairanStore{}
	}, func() rmreq.UpdateRequest[rmreq.CatatanKeseimbanganCairanData] {
		return &rmreq.CatatanKeseimbanganCairanUpdate{}
	}))
	reg("catatan-observasi-bayi", rmctrl.NewController(rmaction.FormCatatanObservasiBayi, func() rmreq.StoreRequest[rmreq.CatatanObservasiBayiData] { return &rmreq.CatatanObservasiBayiStore{} }, func() rmreq.UpdateRequest[rmreq.CatatanObservasiBayiData] { return &rmreq.CatatanObservasiBayiUpdate{} }))
	reg("catatan-observasi-chbp", rmctrl.NewController(rmaction.FormCatatanObservasiChbp, func() rmreq.StoreRequest[rmreq.CatatanObservasiChbpData] { return &rmreq.CatatanObservasiChbpStore{} }, func() rmreq.UpdateRequest[rmreq.CatatanObservasiChbpData] { return &rmreq.CatatanObservasiChbpUpdate{} }))
	reg("catatan-observasi-hemodialisa", rmctrl.NewController(rmaction.FormCatatanObservasiHemodialisa, func() rmreq.StoreRequest[rmreq.CatatanObservasiHemodialisaData] {
		return &rmreq.CatatanObservasiHemodialisaStore{}
	}, func() rmreq.UpdateRequest[rmreq.CatatanObservasiHemodialisaData] {
		return &rmreq.CatatanObservasiHemodialisaUpdate{}
	}))
	reg("catatan-observasi-igd", rmctrl.NewController(rmaction.FormCatatanObservasiIgd, func() rmreq.StoreRequest[rmreq.CatatanObservasiIgdData] { return &rmreq.CatatanObservasiIgdStore{} }, func() rmreq.UpdateRequest[rmreq.CatatanObservasiIgdData] { return &rmreq.CatatanObservasiIgdUpdate{} }))
	reg("catatan-observasi-induksi-persalinan", rmctrl.NewController(rmaction.FormCatatanObservasiInduksiPersalinan, func() rmreq.StoreRequest[rmreq.CatatanObservasiInduksiPersalinanData] {
		return &rmreq.CatatanObservasiInduksiPersalinanStore{}
	}, func() rmreq.UpdateRequest[rmreq.CatatanObservasiInduksiPersalinanData] {
		return &rmreq.CatatanObservasiInduksiPersalinanUpdate{}
	}))
	reg("catatan-observasi-ranap", rmctrl.NewController(rmaction.FormCatatanObservasiRanap, func() rmreq.StoreRequest[rmreq.CatatanObservasiRanapData] { return &rmreq.CatatanObservasiRanapStore{} }, func() rmreq.UpdateRequest[rmreq.CatatanObservasiRanapData] {
		return &rmreq.CatatanObservasiRanapUpdate{}
	}))
	reg("catatan-observasi-ranap-kebidanan", rmctrl.NewController(rmaction.FormCatatanObservasiRanapKebidanan, func() rmreq.StoreRequest[rmreq.CatatanObservasiRanapKebidananData] {
		return &rmreq.CatatanObservasiRanapKebidananStore{}
	}, func() rmreq.UpdateRequest[rmreq.CatatanObservasiRanapKebidananData] {
		return &rmreq.CatatanObservasiRanapKebidananUpdate{}
	}))
	reg("catatan-observasi-ranap-postpartum", rmctrl.NewController(rmaction.FormCatatanObservasiRanapPostpartum, func() rmreq.StoreRequest[rmreq.CatatanObservasiRanapPostpartumData] {
		return &rmreq.CatatanObservasiRanapPostpartumStore{}
	}, func() rmreq.UpdateRequest[rmreq.CatatanObservasiRanapPostpartumData] {
		return &rmreq.CatatanObservasiRanapPostpartumUpdate{}
	}))
	reg("catatan-observasi-restrain-nonfarma", rmctrl.NewController(rmaction.FormCatatanObservasiRestrainNonfarma, func() rmreq.StoreRequest[rmreq.CatatanObservasiRestrainNonfarmaData] {
		return &rmreq.CatatanObservasiRestrainNonfarmaStore{}
	}, func() rmreq.UpdateRequest[rmreq.CatatanObservasiRestrainNonfarmaData] {
		return &rmreq.CatatanObservasiRestrainNonfarmaUpdate{}
	}))
	reg("catatan-observasi-ruang-ok", rmctrl.NewController(rmaction.FormCatatanObservasiRuangOk, func() rmreq.StoreRequest[rmreq.CatatanObservasiRuangOkData] {
		return &rmreq.CatatanObservasiRuangOkStore{}
	}, func() rmreq.UpdateRequest[rmreq.CatatanObservasiRuangOkData] {
		return &rmreq.CatatanObservasiRuangOkUpdate{}
	}))
	reg("catatan-observasi-ventilator", rmctrl.NewController(rmaction.FormCatatanObservasiVentilator, func() rmreq.StoreRequest[rmreq.CatatanObservasiVentilatorData] {
		return &rmreq.CatatanObservasiVentilatorStore{}
	}, func() rmreq.UpdateRequest[rmreq.CatatanObservasiVentilatorData] {
		return &rmreq.CatatanObservasiVentilatorUpdate{}
	}))
	reg("follow-up-dbd", rmctrl.NewController(rmaction.FormFollowUpDbd, func() rmreq.StoreRequest[rmreq.FollowUpDbdData] { return &rmreq.FollowUpDbdStore{} }, func() rmreq.UpdateRequest[rmreq.FollowUpDbdData] { return &rmreq.FollowUpDbdUpdate{} }))
	reg("intervensi-nyeri-farmakologi", rmctrl.NewController(rmaction.FormIntervensiNyeriFarmakologi, func() rmreq.StoreRequest[rmreq.IntervensiNyeriFarmakologiData] {
		return &rmreq.IntervensiNyeriFarmakologiStore{}
	}, func() rmreq.UpdateRequest[rmreq.IntervensiNyeriFarmakologiData] {
		return &rmreq.IntervensiNyeriFarmakologiUpdate{}
	}))
	reg("intervensi-nyeri-nonfarmakologi", rmctrl.NewController(rmaction.FormIntervensiNyeriNonfarmakologi, func() rmreq.StoreRequest[rmreq.IntervensiNyeriNonfarmakologiData] {
		return &rmreq.IntervensiNyeriNonfarmakologiStore{}
	}, func() rmreq.UpdateRequest[rmreq.IntervensiNyeriNonfarmakologiData] {
		return &rmreq.IntervensiNyeriNonfarmakologiUpdate{}
	}))
	reg("monitoring-asuhan-gizi", rmctrl.NewController(rmaction.FormMonitoringAsuhanGizi, func() rmreq.StoreRequest[rmreq.MonitoringAsuhanGiziData] { return &rmreq.MonitoringAsuhanGiziStore{} }, func() rmreq.UpdateRequest[rmreq.MonitoringAsuhanGiziData] { return &rmreq.MonitoringAsuhanGiziUpdate{} }))
	reg("monitoring-reaksi-tranfusi", rmctrl.NewController(rmaction.FormMonitoringReaksiTranfusi, func() rmreq.StoreRequest[rmreq.MonitoringReaksiTranfusiData] {
		return &rmreq.MonitoringReaksiTranfusiStore{}
	}, func() rmreq.UpdateRequest[rmreq.MonitoringReaksiTranfusiData] {
		return &rmreq.MonitoringReaksiTranfusiUpdate{}
	}))
	reg("resume-pasien", rmctrl.NewController(rmaction.FormResumePasien, func() rmreq.StoreRequest[rmreq.ResumePasienData] { return &rmreq.ResumePasienStore{} }, func() rmreq.UpdateRequest[rmreq.ResumePasienData] { return &rmreq.ResumePasienUpdate{} }))
	reg("resume-pasien-ranap", rmctrl.NewController(rmaction.FormResumePasienRanap, func() rmreq.StoreRequest[rmreq.ResumePasienRanapData] { return &rmreq.ResumePasienRanapStore{} }, func() rmreq.UpdateRequest[rmreq.ResumePasienRanapData] { return &rmreq.ResumePasienRanapUpdate{} }))
	reg("skrining-gizi-kehamilan", rmctrl.NewController(rmaction.FormSkriningGiziKehamilan, func() rmreq.StoreRequest[rmreq.SkriningGiziKehamilanData] { return &rmreq.SkriningGiziKehamilanStore{} }, func() rmreq.UpdateRequest[rmreq.SkriningGiziKehamilanData] {
		return &rmreq.SkriningGiziKehamilanUpdate{}
	}))
	reg("skrining-gizi", rmctrl.NewController(rmaction.FormSkriningGizi, func() rmreq.StoreRequest[rmreq.SkriningGiziData] { return &rmreq.SkriningGiziStore{} }, func() rmreq.UpdateRequest[rmreq.SkriningGiziData] { return &rmreq.SkriningGiziUpdate{} }))
	reg("deteksi-dini-corona", rmctrl.NewController(rmaction.FormDeteksiDiniCorona, func() rmreq.StoreRequest[rmreq.DeteksiDiniCoronaData] { return &rmreq.DeteksiDiniCoronaStore{} }, func() rmreq.UpdateRequest[rmreq.DeteksiDiniCoronaData] { return &rmreq.DeteksiDiniCoronaUpdate{} }))
	reg("edukasi-pasien-keluarga-rj", rmctrl.NewController(rmaction.FormEdukasiPasienKeluargaRj, func() rmreq.StoreRequest[rmreq.EdukasiPasienKeluargaRjData] {
		return &rmreq.EdukasiPasienKeluargaRjStore{}
	}, func() rmreq.UpdateRequest[rmreq.EdukasiPasienKeluargaRjData] {
		return &rmreq.EdukasiPasienKeluargaRjUpdate{}
	}))
	reg("hasil-endoskopi-faring-laring", rmctrl.NewController(rmaction.FormHasilEndoskopiFaringLaring, func() rmreq.StoreRequest[rmreq.HasilEndoskopiFaringLaringData] {
		return &rmreq.HasilEndoskopiFaringLaringStore{}
	}, func() rmreq.UpdateRequest[rmreq.HasilEndoskopiFaringLaringData] {
		return &rmreq.HasilEndoskopiFaringLaringUpdate{}
	}))
	reg("hasil-endoskopi-hidung", rmctrl.NewController(rmaction.FormHasilEndoskopiHidung, func() rmreq.StoreRequest[rmreq.HasilEndoskopiHidungData] { return &rmreq.HasilEndoskopiHidungStore{} }, func() rmreq.UpdateRequest[rmreq.HasilEndoskopiHidungData] { return &rmreq.HasilEndoskopiHidungUpdate{} }))
	reg("hasil-endoskopi-telinga", rmctrl.NewController(rmaction.FormHasilEndoskopiTelinga, func() rmreq.StoreRequest[rmreq.HasilEndoskopiTelingaData] { return &rmreq.HasilEndoskopiTelingaStore{} }, func() rmreq.UpdateRequest[rmreq.HasilEndoskopiTelingaData] {
		return &rmreq.HasilEndoskopiTelingaUpdate{}
	}))
	reg("hasil-pemeriksaan-ekg", rmctrl.NewController(rmaction.FormHasilPemeriksaanEkg, func() rmreq.StoreRequest[rmreq.HasilPemeriksaanEkgData] { return &rmreq.HasilPemeriksaanEkgStore{} }, func() rmreq.UpdateRequest[rmreq.HasilPemeriksaanEkgData] { return &rmreq.HasilPemeriksaanEkgUpdate{} }))
	reg("hasil-pemeriksaan-echo", rmctrl.NewController(rmaction.FormHasilPemeriksaanEcho, func() rmreq.StoreRequest[rmreq.HasilPemeriksaanEchoData] { return &rmreq.HasilPemeriksaanEchoStore{} }, func() rmreq.UpdateRequest[rmreq.HasilPemeriksaanEchoData] { return &rmreq.HasilPemeriksaanEchoUpdate{} }))
	reg("hasil-pemeriksaan-echo-pediatrik", rmctrl.NewController(rmaction.FormHasilPemeriksaanEchoPediatrik, func() rmreq.StoreRequest[rmreq.HasilPemeriksaanEchoPediatrikData] {
		return &rmreq.HasilPemeriksaanEchoPediatrikStore{}
	}, func() rmreq.UpdateRequest[rmreq.HasilPemeriksaanEchoPediatrikData] {
		return &rmreq.HasilPemeriksaanEchoPediatrikUpdate{}
	}))
	reg("hasil-pemeriksaan-oct", rmctrl.NewController(rmaction.FormHasilPemeriksaanOct, func() rmreq.StoreRequest[rmreq.HasilPemeriksaanOctData] { return &rmreq.HasilPemeriksaanOctStore{} }, func() rmreq.UpdateRequest[rmreq.HasilPemeriksaanOctData] { return &rmreq.HasilPemeriksaanOctUpdate{} }))
	reg("hasil-pemeriksaan-slit-lamp", rmctrl.NewController(rmaction.FormHasilPemeriksaanSlitLamp, func() rmreq.StoreRequest[rmreq.HasilPemeriksaanSlitLampData] {
		return &rmreq.HasilPemeriksaanSlitLampStore{}
	}, func() rmreq.UpdateRequest[rmreq.HasilPemeriksaanSlitLampData] {
		return &rmreq.HasilPemeriksaanSlitLampUpdate{}
	}))
	reg("hasil-pemeriksaan-treadmill", rmctrl.NewController(rmaction.FormHasilPemeriksaanTreadmill, func() rmreq.StoreRequest[rmreq.HasilPemeriksaanTreadmillData] {
		return &rmreq.HasilPemeriksaanTreadmillStore{}
	}, func() rmreq.UpdateRequest[rmreq.HasilPemeriksaanTreadmillData] {
		return &rmreq.HasilPemeriksaanTreadmillUpdate{}
	}))
	reg("hasil-pemeriksaan-usg", rmctrl.NewController(rmaction.FormHasilPemeriksaanUsg, func() rmreq.StoreRequest[rmreq.HasilPemeriksaanUsgData] { return &rmreq.HasilPemeriksaanUsgStore{} }, func() rmreq.UpdateRequest[rmreq.HasilPemeriksaanUsgData] { return &rmreq.HasilPemeriksaanUsgUpdate{} }))
	reg("hasil-pemeriksaan-usg-abdomen", rmctrl.NewController(rmaction.FormHasilPemeriksaanUsgAbdomen, func() rmreq.StoreRequest[rmreq.HasilPemeriksaanUsgAbdomenData] {
		return &rmreq.HasilPemeriksaanUsgAbdomenStore{}
	}, func() rmreq.UpdateRequest[rmreq.HasilPemeriksaanUsgAbdomenData] {
		return &rmreq.HasilPemeriksaanUsgAbdomenUpdate{}
	}))
	reg("hasil-pemeriksaan-usg-gynecologi", rmctrl.NewController(rmaction.FormHasilPemeriksaanUsgGynecologi, func() rmreq.StoreRequest[rmreq.HasilPemeriksaanUsgGynecologiData] {
		return &rmreq.HasilPemeriksaanUsgGynecologiStore{}
	}, func() rmreq.UpdateRequest[rmreq.HasilPemeriksaanUsgGynecologiData] {
		return &rmreq.HasilPemeriksaanUsgGynecologiUpdate{}
	}))
	reg("hasil-pemeriksaan-usg-neonatus", rmctrl.NewController(rmaction.FormHasilPemeriksaanUsgNeonatus, func() rmreq.StoreRequest[rmreq.HasilPemeriksaanUsgNeonatusData] {
		return &rmreq.HasilPemeriksaanUsgNeonatusStore{}
	}, func() rmreq.UpdateRequest[rmreq.HasilPemeriksaanUsgNeonatusData] {
		return &rmreq.HasilPemeriksaanUsgNeonatusUpdate{}
	}))
	reg("hasil-pemeriksaan-usg-urologi", rmctrl.NewController(rmaction.FormHasilPemeriksaanUsgUrologi, func() rmreq.StoreRequest[rmreq.HasilPemeriksaanUsgUrologiData] {
		return &rmreq.HasilPemeriksaanUsgUrologiStore{}
	}, func() rmreq.UpdateRequest[rmreq.HasilPemeriksaanUsgUrologiData] {
		return &rmreq.HasilPemeriksaanUsgUrologiUpdate{}
	}))
	reg("hasil-tindakan-eswl", rmctrl.NewController(rmaction.FormHasilTindakanEswl, func() rmreq.StoreRequest[rmreq.HasilTindakanEswlData] { return &rmreq.HasilTindakanEswlStore{} }, func() rmreq.UpdateRequest[rmreq.HasilTindakanEswlData] { return &rmreq.HasilTindakanEswlUpdate{} }))
	reg("hemodialisa", rmctrl.NewController(rmaction.FormHemodialisa, func() rmreq.StoreRequest[rmreq.HemodialisaData] { return &rmreq.HemodialisaStore{} }, func() rmreq.UpdateRequest[rmreq.HemodialisaData] { return &rmreq.HemodialisaUpdate{} }))
	reg("konseling-farmasi", rmctrl.NewController(rmaction.FormKonselingFarmasi, func() rmreq.StoreRequest[rmreq.KonselingFarmasiData] { return &rmreq.KonselingFarmasiStore{} }, func() rmreq.UpdateRequest[rmreq.KonselingFarmasiData] { return &rmreq.KonselingFarmasiUpdate{} }))
	reg("penilaian-awal-keperawatan-kebidanan-ranap", rmctrl.NewController(rmaction.FormPenilaianAwalKeperawatanKebidananRanap, func() rmreq.StoreRequest[rmreq.PenilaianAwalKeperawatanKebidananRanapData] {
		return &rmreq.PenilaianAwalKeperawatanKebidananRanapStore{}
	}, func() rmreq.UpdateRequest[rmreq.PenilaianAwalKeperawatanKebidananRanapData] {
		return &rmreq.PenilaianAwalKeperawatanKebidananRanapUpdate{}
	}))
	reg("laporan-tindakan", rmctrl.NewController(rmaction.FormLaporanTindakan, func() rmreq.StoreRequest[rmreq.LaporanTindakanData] { return &rmreq.LaporanTindakanStore{} }, func() rmreq.UpdateRequest[rmreq.LaporanTindakanData] { return &rmreq.LaporanTindakanUpdate{} }))
	reg("layanan-kedokteran-fisik-rehabilitasi", rmctrl.NewController(rmaction.FormLayananKedokteranFisikRehabilitasi, func() rmreq.StoreRequest[rmreq.LayananKedokteranFisikRehabilitasiData] {
		return &rmreq.LayananKedokteranFisikRehabilitasiStore{}
	}, func() rmreq.UpdateRequest[rmreq.LayananKedokteranFisikRehabilitasiData] {
		return &rmreq.LayananKedokteranFisikRehabilitasiUpdate{}
	}))
	reg("layanan-program-kfr", rmctrl.NewController(rmaction.FormLayananProgramKfr, func() rmreq.StoreRequest[rmreq.LayananProgramKfrData] { return &rmreq.LayananProgramKfrStore{} }, func() rmreq.UpdateRequest[rmreq.LayananProgramKfrData] { return &rmreq.LayananProgramKfrUpdate{} }))
	reg("penilaian-mcu", rmctrl.NewController(rmaction.FormPenilaianMcu, func() rmreq.StoreRequest[rmreq.PenilaianMcuData] { return &rmreq.PenilaianMcuStore{} }, func() rmreq.UpdateRequest[rmreq.PenilaianMcuData] { return &rmreq.PenilaianMcuUpdate{} }))
	reg("skor-aldrette-pasca-anestesi", rmctrl.NewController(rmaction.FormSkorAldrettePascaAnestesi, func() rmreq.StoreRequest[rmreq.SkorAldrettePascaAnestesiData] {
		return &rmreq.SkorAldrettePascaAnestesiStore{}
	}, func() rmreq.UpdateRequest[rmreq.SkorAldrettePascaAnestesiData] {
		return &rmreq.SkorAldrettePascaAnestesiUpdate{}
	}))
	reg("skor-bromage-pasca-anestesi", rmctrl.NewController(rmaction.FormSkorBromagePascaAnestesi, func() rmreq.StoreRequest[rmreq.SkorBromagePascaAnestesiData] {
		return &rmreq.SkorBromagePascaAnestesiStore{}
	}, func() rmreq.UpdateRequest[rmreq.SkorBromagePascaAnestesiData] {
		return &rmreq.SkorBromagePascaAnestesiUpdate{}
	}))
	reg("skor-steward-pasca-anestesi", rmctrl.NewController(rmaction.FormSkorStewardPascaAnestesi, func() rmreq.StoreRequest[rmreq.SkorStewardPascaAnestesiData] {
		return &rmreq.SkorStewardPascaAnestesiStore{}
	}, func() rmreq.UpdateRequest[rmreq.SkorStewardPascaAnestesiData] {
		return &rmreq.SkorStewardPascaAnestesiUpdate{}
	}))
	reg("pelaksanaan-informasi-edukasi", rmctrl.NewController(rmaction.FormPelaksanaanInformasiEdukasi, func() rmreq.StoreRequest[rmreq.PelaksanaanInformasiEdukasiData] {
		return &rmreq.PelaksanaanInformasiEdukasiStore{}
	}, func() rmreq.UpdateRequest[rmreq.PelaksanaanInformasiEdukasiData] {
		return &rmreq.PelaksanaanInformasiEdukasiUpdate{}
	}))
	reg("monitoring-efek-samping-obat", rmctrl.NewController(rmaction.FormMonitoringEfekSampingObat, func() rmreq.StoreRequest[rmreq.MonitoringEfekSampingObatData] {
		return &rmreq.MonitoringEfekSampingObatStore{}
	}, func() rmreq.UpdateRequest[rmreq.MonitoringEfekSampingObatData] {
		return &rmreq.MonitoringEfekSampingObatUpdate{}
	}))
	reg("pemantauan-pews-dewasa", rmctrl.NewController(rmaction.FormPemantauanPewsDewasa, func() rmreq.StoreRequest[rmreq.PemantauanPewsDewasaData] { return &rmreq.PemantauanPewsDewasaStore{} }, func() rmreq.UpdateRequest[rmreq.PemantauanPewsDewasaData] { return &rmreq.PemantauanPewsDewasaUpdate{} }))
	reg("pemantauan-ews-neonatus", rmctrl.NewController(rmaction.FormPemantauanEwsNeonatus, func() rmreq.StoreRequest[rmreq.PemantauanEwsNeonatusData] { return &rmreq.PemantauanEwsNeonatusStore{} }, func() rmreq.UpdateRequest[rmreq.PemantauanEwsNeonatusData] {
		return &rmreq.PemantauanEwsNeonatusUpdate{}
	}))
	reg("pemantauan-meows-obstetri", rmctrl.NewController(rmaction.FormPemantauanMeowsObstetri, func() rmreq.StoreRequest[rmreq.PemantauanMeowsObstetriData] {
		return &rmreq.PemantauanMeowsObstetriStore{}
	}, func() rmreq.UpdateRequest[rmreq.PemantauanMeowsObstetriData] {
		return &rmreq.PemantauanMeowsObstetriUpdate{}
	}))
	reg("pemantauan-pews-anak", rmctrl.NewController(rmaction.FormPemantauanPewsAnak, func() rmreq.StoreRequest[rmreq.PemantauanPewsAnakData] { return &rmreq.PemantauanPewsAnakStore{} }, func() rmreq.UpdateRequest[rmreq.PemantauanPewsAnakData] { return &rmreq.PemantauanPewsAnakUpdate{} }))
	reg("penatalaksanaan-terapi-okupasi", rmctrl.NewController(rmaction.FormPenatalaksanaanTerapiOkupasi, func() rmreq.StoreRequest[rmreq.PenatalaksanaanTerapiOkupasiData] {
		return &rmreq.PenatalaksanaanTerapiOkupasiStore{}
	}, func() rmreq.UpdateRequest[rmreq.PenatalaksanaanTerapiOkupasiData] {
		return &rmreq.PenatalaksanaanTerapiOkupasiUpdate{}
	}))
	reg("pengkajian-restrain", rmctrl.NewController(rmaction.FormPengkajianRestrain, func() rmreq.StoreRequest[rmreq.PengkajianRestrainData] { return &rmreq.PengkajianRestrainStore{} }, func() rmreq.UpdateRequest[rmreq.PengkajianRestrainData] { return &rmreq.PengkajianRestrainUpdate{} }))
	reg("penilaian-awal-keperawatan-ralan-bayi", rmctrl.NewController(rmaction.FormPenilaianAwalKeperawatanRalanBayi, func() rmreq.StoreRequest[rmreq.PenilaianAwalKeperawatanRalanBayiData] {
		return &rmreq.PenilaianAwalKeperawatanRalanBayiStore{}
	}, func() rmreq.UpdateRequest[rmreq.PenilaianAwalKeperawatanRalanBayiData] {
		return &rmreq.PenilaianAwalKeperawatanRalanBayiUpdate{}
	}))
	reg("penilaian-awal-keperawatan-gigi", rmctrl.NewController(rmaction.FormPenilaianAwalKeperawatanGigi, func() rmreq.StoreRequest[rmreq.PenilaianAwalKeperawatanGigiData] {
		return &rmreq.PenilaianAwalKeperawatanGigiStore{}
	}, func() rmreq.UpdateRequest[rmreq.PenilaianAwalKeperawatanGigiData] {
		return &rmreq.PenilaianAwalKeperawatanGigiUpdate{}
	}))
	reg("penilaian-awal-keperawatan-igd", rmctrl.NewController(rmaction.FormPenilaianAwalKeperawatanIgd, func() rmreq.StoreRequest[rmreq.PenilaianAwalKeperawatanIgdData] {
		return &rmreq.PenilaianAwalKeperawatanIgdStore{}
	}, func() rmreq.UpdateRequest[rmreq.PenilaianAwalKeperawatanIgdData] {
		return &rmreq.PenilaianAwalKeperawatanIgdUpdate{}
	}))
	reg("penilaian-awal-keperawatan-kebidanan", rmctrl.NewController(rmaction.FormPenilaianAwalKeperawatanKebidanan, func() rmreq.StoreRequest[rmreq.PenilaianAwalKeperawatanKebidananData] {
		return &rmreq.PenilaianAwalKeperawatanKebidananStore{}
	}, func() rmreq.UpdateRequest[rmreq.PenilaianAwalKeperawatanKebidananData] {
		return &rmreq.PenilaianAwalKeperawatanKebidananUpdate{}
	}))
	reg("penilaian-awal-keperawatan-mata", rmctrl.NewController(rmaction.FormPenilaianAwalKeperawatanMata, func() rmreq.StoreRequest[rmreq.PenilaianAwalKeperawatanMataData] {
		return &rmreq.PenilaianAwalKeperawatanMataStore{}
	}, func() rmreq.UpdateRequest[rmreq.PenilaianAwalKeperawatanMataData] {
		return &rmreq.PenilaianAwalKeperawatanMataUpdate{}
	}))
	reg("penilaian-awal-keperawatan-ralan", rmctrl.NewController(rmaction.FormPenilaianAwalKeperawatanRalan, func() rmreq.StoreRequest[rmreq.PenilaianAwalKeperawatanRalanData] {
		return &rmreq.PenilaianAwalKeperawatanRalanStore{}
	}, func() rmreq.UpdateRequest[rmreq.PenilaianAwalKeperawatanRalanData] {
		return &rmreq.PenilaianAwalKeperawatanRalanUpdate{}
	}))
	reg("penilaian-awal-keperawatan-ralan-geriatri", rmctrl.NewController(rmaction.FormPenilaianAwalKeperawatanRalanGeriatri, func() rmreq.StoreRequest[rmreq.PenilaianAwalKeperawatanRalanGeriatriData] {
		return &rmreq.PenilaianAwalKeperawatanRalanGeriatriStore{}
	}, func() rmreq.UpdateRequest[rmreq.PenilaianAwalKeperawatanRalanGeriatriData] {
		return &rmreq.PenilaianAwalKeperawatanRalanGeriatriUpdate{}
	}))
	reg("penilaian-awal-keperawatan-ralan-psikiatri", rmctrl.NewController(rmaction.FormPenilaianAwalKeperawatanRalanPsikiatri, func() rmreq.StoreRequest[rmreq.PenilaianAwalKeperawatanRalanPsikiatriData] {
		return &rmreq.PenilaianAwalKeperawatanRalanPsikiatriStore{}
	}, func() rmreq.UpdateRequest[rmreq.PenilaianAwalKeperawatanRalanPsikiatriData] {
		return &rmreq.PenilaianAwalKeperawatanRalanPsikiatriUpdate{}
	}))
	reg("penilaian-awal-keperawatan-ranap", rmctrl.NewController(rmaction.FormPenilaianAwalKeperawatanRanap, func() rmreq.StoreRequest[rmreq.PenilaianAwalKeperawatanRanapData] {
		return &rmreq.PenilaianAwalKeperawatanRanapStore{}
	}, func() rmreq.UpdateRequest[rmreq.PenilaianAwalKeperawatanRanapData] {
		return &rmreq.PenilaianAwalKeperawatanRanapUpdate{}
	}))
	reg("penilaian-awal-keperawatan-ranap-bayi", rmctrl.NewController(rmaction.FormPenilaianAwalKeperawatanRanapBayi, func() rmreq.StoreRequest[rmreq.PenilaianAwalKeperawatanRanapBayiData] {
		return &rmreq.PenilaianAwalKeperawatanRanapBayiStore{}
	}, func() rmreq.UpdateRequest[rmreq.PenilaianAwalKeperawatanRanapBayiData] {
		return &rmreq.PenilaianAwalKeperawatanRanapBayiUpdate{}
	}))
	reg("penilaian-awal-keperawatan-ranap-neonatus", rmctrl.NewController(rmaction.FormPenilaianAwalKeperawatanRanapNeonatus, func() rmreq.StoreRequest[rmreq.PenilaianAwalKeperawatanRanapNeonatusData] {
		return &rmreq.PenilaianAwalKeperawatanRanapNeonatusStore{}
	}, func() rmreq.UpdateRequest[rmreq.PenilaianAwalKeperawatanRanapNeonatusData] {
		return &rmreq.PenilaianAwalKeperawatanRanapNeonatusUpdate{}
	}))
	reg("penilaian-medis-hemodialisa", rmctrl.NewController(rmaction.FormPenilaianMedisHemodialisa, func() rmreq.StoreRequest[rmreq.PenilaianMedisHemodialisaData] {
		return &rmreq.PenilaianMedisHemodialisaStore{}
	}, func() rmreq.UpdateRequest[rmreq.PenilaianMedisHemodialisaData] {
		return &rmreq.PenilaianMedisHemodialisaUpdate{}
	}))
	reg("penilaian-medis-igd", rmctrl.NewController(rmaction.FormPenilaianMedisIgd, func() rmreq.StoreRequest[rmreq.PenilaianMedisIgdData] { return &rmreq.PenilaianMedisIgdStore{} }, func() rmreq.UpdateRequest[rmreq.PenilaianMedisIgdData] { return &rmreq.PenilaianMedisIgdUpdate{} }))
	reg("penilaian-medis-ralan-gawat-darurat-psikiatri", rmctrl.NewController(rmaction.FormPenilaianMedisRalanGawatDaruratPsikiatri, func() rmreq.StoreRequest[rmreq.PenilaianMedisRalanGawatDaruratPsikiatriData] {
		return &rmreq.PenilaianMedisRalanGawatDaruratPsikiatriStore{}
	}, func() rmreq.UpdateRequest[rmreq.PenilaianMedisRalanGawatDaruratPsikiatriData] {
		return &rmreq.PenilaianMedisRalanGawatDaruratPsikiatriUpdate{}
	}))
	reg("penilaian-medis-ralan-anak", rmctrl.NewController(rmaction.FormPenilaianMedisRalanAnak, func() rmreq.StoreRequest[rmreq.PenilaianMedisRalanAnakData] {
		return &rmreq.PenilaianMedisRalanAnakStore{}
	}, func() rmreq.UpdateRequest[rmreq.PenilaianMedisRalanAnakData] {
		return &rmreq.PenilaianMedisRalanAnakUpdate{}
	}))
	reg("penilaian-medis-ralan-bedah", rmctrl.NewController(rmaction.FormPenilaianMedisRalanBedah, func() rmreq.StoreRequest[rmreq.PenilaianMedisRalanBedahData] {
		return &rmreq.PenilaianMedisRalanBedahStore{}
	}, func() rmreq.UpdateRequest[rmreq.PenilaianMedisRalanBedahData] {
		return &rmreq.PenilaianMedisRalanBedahUpdate{}
	}))
	reg("penilaian-medis-ralan-bedah-mulut", rmctrl.NewController(rmaction.FormPenilaianMedisRalanBedahMulut, func() rmreq.StoreRequest[rmreq.PenilaianMedisRalanBedahMulutData] {
		return &rmreq.PenilaianMedisRalanBedahMulutStore{}
	}, func() rmreq.UpdateRequest[rmreq.PenilaianMedisRalanBedahMulutData] {
		return &rmreq.PenilaianMedisRalanBedahMulutUpdate{}
	}))
	reg("penilaian-medis-ralan", rmctrl.NewController(rmaction.FormPenilaianMedisRalan, func() rmreq.StoreRequest[rmreq.PenilaianMedisRalanData] { return &rmreq.PenilaianMedisRalanStore{} }, func() rmreq.UpdateRequest[rmreq.PenilaianMedisRalanData] { return &rmreq.PenilaianMedisRalanUpdate{} }))
	reg("penilaian-medis-ralan-geriatri", rmctrl.NewController(rmaction.FormPenilaianMedisRalanGeriatri, func() rmreq.StoreRequest[rmreq.PenilaianMedisRalanGeriatriData] {
		return &rmreq.PenilaianMedisRalanGeriatriStore{}
	}, func() rmreq.UpdateRequest[rmreq.PenilaianMedisRalanGeriatriData] {
		return &rmreq.PenilaianMedisRalanGeriatriUpdate{}
	}))
	reg("penilaian-medis-ralan-jantung", rmctrl.NewController(rmaction.FormPenilaianMedisRalanJantung, func() rmreq.StoreRequest[rmreq.PenilaianMedisRalanJantungData] {
		return &rmreq.PenilaianMedisRalanJantungStore{}
	}, func() rmreq.UpdateRequest[rmreq.PenilaianMedisRalanJantungData] {
		return &rmreq.PenilaianMedisRalanJantungUpdate{}
	}))
	reg("penilaian-medis-ralan-kandungan", rmctrl.NewController(rmaction.FormPenilaianMedisRalanKandungan, func() rmreq.StoreRequest[rmreq.PenilaianMedisRalanKandunganData] {
		return &rmreq.PenilaianMedisRalanKandunganStore{}
	}, func() rmreq.UpdateRequest[rmreq.PenilaianMedisRalanKandunganData] {
		return &rmreq.PenilaianMedisRalanKandunganUpdate{}
	}))
	reg("penilaian-medis-ralan-kulitdankelamin", rmctrl.NewController(rmaction.FormPenilaianMedisRalanKulitdankelamin, func() rmreq.StoreRequest[rmreq.PenilaianMedisRalanKulitdankelaminData] {
		return &rmreq.PenilaianMedisRalanKulitdankelaminStore{}
	}, func() rmreq.UpdateRequest[rmreq.PenilaianMedisRalanKulitdankelaminData] {
		return &rmreq.PenilaianMedisRalanKulitdankelaminUpdate{}
	}))
	reg("penilaian-medis-ralan-mata", rmctrl.NewController(rmaction.FormPenilaianMedisRalanMata, func() rmreq.StoreRequest[rmreq.PenilaianMedisRalanMataData] {
		return &rmreq.PenilaianMedisRalanMataStore{}
	}, func() rmreq.UpdateRequest[rmreq.PenilaianMedisRalanMataData] {
		return &rmreq.PenilaianMedisRalanMataUpdate{}
	}))
	reg("penilaian-medis-ralan-neurologi", rmctrl.NewController(rmaction.FormPenilaianMedisRalanNeurologi, func() rmreq.StoreRequest[rmreq.PenilaianMedisRalanNeurologiData] {
		return &rmreq.PenilaianMedisRalanNeurologiStore{}
	}, func() rmreq.UpdateRequest[rmreq.PenilaianMedisRalanNeurologiData] {
		return &rmreq.PenilaianMedisRalanNeurologiUpdate{}
	}))
	reg("penilaian-medis-ralan-orthopedi", rmctrl.NewController(rmaction.FormPenilaianMedisRalanOrthopedi, func() rmreq.StoreRequest[rmreq.PenilaianMedisRalanOrthopediData] {
		return &rmreq.PenilaianMedisRalanOrthopediStore{}
	}, func() rmreq.UpdateRequest[rmreq.PenilaianMedisRalanOrthopediData] {
		return &rmreq.PenilaianMedisRalanOrthopediUpdate{}
	}))
	reg("penilaian-medis-ralan-paru", rmctrl.NewController(rmaction.FormPenilaianMedisRalanParu, func() rmreq.StoreRequest[rmreq.PenilaianMedisRalanParuData] {
		return &rmreq.PenilaianMedisRalanParuStore{}
	}, func() rmreq.UpdateRequest[rmreq.PenilaianMedisRalanParuData] {
		return &rmreq.PenilaianMedisRalanParuUpdate{}
	}))
	reg("penilaian-medis-ralan-penyakit-dalam", rmctrl.NewController(rmaction.FormPenilaianMedisRalanPenyakitDalam, func() rmreq.StoreRequest[rmreq.PenilaianMedisRalanPenyakitDalamData] {
		return &rmreq.PenilaianMedisRalanPenyakitDalamStore{}
	}, func() rmreq.UpdateRequest[rmreq.PenilaianMedisRalanPenyakitDalamData] {
		return &rmreq.PenilaianMedisRalanPenyakitDalamUpdate{}
	}))
	reg("penilaian-medis-ralan-psikiatrik", rmctrl.NewController(rmaction.FormPenilaianMedisRalanPsikiatrik, func() rmreq.StoreRequest[rmreq.PenilaianMedisRalanPsikiatrikData] {
		return &rmreq.PenilaianMedisRalanPsikiatrikStore{}
	}, func() rmreq.UpdateRequest[rmreq.PenilaianMedisRalanPsikiatrikData] {
		return &rmreq.PenilaianMedisRalanPsikiatrikUpdate{}
	}))
	reg("penilaian-medis-ralan-rehab-medik", rmctrl.NewController(rmaction.FormPenilaianMedisRalanRehabMedik, func() rmreq.StoreRequest[rmreq.PenilaianMedisRalanRehabMedikData] {
		return &rmreq.PenilaianMedisRalanRehabMedikStore{}
	}, func() rmreq.UpdateRequest[rmreq.PenilaianMedisRalanRehabMedikData] {
		return &rmreq.PenilaianMedisRalanRehabMedikUpdate{}
	}))
	reg("penilaian-medis-ralan-tht", rmctrl.NewController(rmaction.FormPenilaianMedisRalanTht, func() rmreq.StoreRequest[rmreq.PenilaianMedisRalanThtData] {
		return &rmreq.PenilaianMedisRalanThtStore{}
	}, func() rmreq.UpdateRequest[rmreq.PenilaianMedisRalanThtData] {
		return &rmreq.PenilaianMedisRalanThtUpdate{}
	}))
	reg("penilaian-medis-ralan-urologi", rmctrl.NewController(rmaction.FormPenilaianMedisRalanUrologi, func() rmreq.StoreRequest[rmreq.PenilaianMedisRalanUrologiData] {
		return &rmreq.PenilaianMedisRalanUrologiStore{}
	}, func() rmreq.UpdateRequest[rmreq.PenilaianMedisRalanUrologiData] {
		return &rmreq.PenilaianMedisRalanUrologiUpdate{}
	}))
	reg("penilaian-medis-ranap", rmctrl.NewController(rmaction.FormPenilaianMedisRanap, func() rmreq.StoreRequest[rmreq.PenilaianMedisRanapData] { return &rmreq.PenilaianMedisRanapStore{} }, func() rmreq.UpdateRequest[rmreq.PenilaianMedisRanapData] { return &rmreq.PenilaianMedisRanapUpdate{} }))
	reg("penilaian-medis-ranap-jantung", rmctrl.NewController(rmaction.FormPenilaianMedisRanapJantung, func() rmreq.StoreRequest[rmreq.PenilaianMedisRanapJantungData] {
		return &rmreq.PenilaianMedisRanapJantungStore{}
	}, func() rmreq.UpdateRequest[rmreq.PenilaianMedisRanapJantungData] {
		return &rmreq.PenilaianMedisRanapJantungUpdate{}
	}))
	reg("penilaian-medis-ranap-kandungan", rmctrl.NewController(rmaction.FormPenilaianMedisRanapKandungan, func() rmreq.StoreRequest[rmreq.PenilaianMedisRanapKandunganData] {
		return &rmreq.PenilaianMedisRanapKandunganStore{}
	}, func() rmreq.UpdateRequest[rmreq.PenilaianMedisRanapKandunganData] {
		return &rmreq.PenilaianMedisRanapKandunganUpdate{}
	}))
	reg("penilaian-medis-ranap-neonatus", rmctrl.NewController(rmaction.FormPenilaianMedisRanapNeonatus, func() rmreq.StoreRequest[rmreq.PenilaianMedisRanapNeonatusData] {
		return &rmreq.PenilaianMedisRanapNeonatusStore{}
	}, func() rmreq.UpdateRequest[rmreq.PenilaianMedisRanapNeonatusData] {
		return &rmreq.PenilaianMedisRanapNeonatusUpdate{}
	}))
	reg("penilaian-medis-ranap-psikiatrik", rmctrl.NewController(rmaction.FormPenilaianMedisRanapPsikiatrik, func() rmreq.StoreRequest[rmreq.PenilaianMedisRanapPsikiatrikData] {
		return &rmreq.PenilaianMedisRanapPsikiatrikStore{}
	}, func() rmreq.UpdateRequest[rmreq.PenilaianMedisRanapPsikiatrikData] {
		return &rmreq.PenilaianMedisRanapPsikiatrikUpdate{}
	}))
	reg("penilaian-bayi-baru-lahir", rmctrl.NewController(rmaction.FormPenilaianBayiBaruLahir, func() rmreq.StoreRequest[rmreq.PenilaianBayiBaruLahirData] {
		return &rmreq.PenilaianBayiBaruLahirStore{}
	}, func() rmreq.UpdateRequest[rmreq.PenilaianBayiBaruLahirData] {
		return &rmreq.PenilaianBayiBaruLahirUpdate{}
	}))
	reg("penilaian-dehidrasi", rmctrl.NewController(rmaction.FormPenilaianDehidrasi, func() rmreq.StoreRequest[rmreq.PenilaianDehidrasiData] { return &rmreq.PenilaianDehidrasiStore{} }, func() rmreq.UpdateRequest[rmreq.PenilaianDehidrasiData] { return &rmreq.PenilaianDehidrasiUpdate{} }))
	reg("penilaian-fisioterapi", rmctrl.NewController(rmaction.FormPenilaianFisioterapi, func() rmreq.StoreRequest[rmreq.PenilaianFisioterapiData] { return &rmreq.PenilaianFisioterapiStore{} }, func() rmreq.UpdateRequest[rmreq.PenilaianFisioterapiData] { return &rmreq.PenilaianFisioterapiUpdate{} }))
	reg("penilaian-korban-kekerasan", rmctrl.NewController(rmaction.FormPenilaianKorbanKekerasan, func() rmreq.StoreRequest[rmreq.PenilaianKorbanKekerasanData] {
		return &rmreq.PenilaianKorbanKekerasanStore{}
	}, func() rmreq.UpdateRequest[rmreq.PenilaianKorbanKekerasanData] {
		return &rmreq.PenilaianKorbanKekerasanUpdate{}
	}))
	reg("penilaian-lanjutan-resiko-jatuh-anak", rmctrl.NewController(rmaction.FormPenilaianLanjutanResikoJatuhAnak, func() rmreq.StoreRequest[rmreq.PenilaianLanjutanResikoJatuhAnakData] {
		return &rmreq.PenilaianLanjutanResikoJatuhAnakStore{}
	}, func() rmreq.UpdateRequest[rmreq.PenilaianLanjutanResikoJatuhAnakData] {
		return &rmreq.PenilaianLanjutanResikoJatuhAnakUpdate{}
	}))
	reg("penilaian-lanjutan-resiko-jatuh-dewasa", rmctrl.NewController(rmaction.FormPenilaianLanjutanResikoJatuhDewasa, func() rmreq.StoreRequest[rmreq.PenilaianLanjutanResikoJatuhDewasaData] {
		return &rmreq.PenilaianLanjutanResikoJatuhDewasaStore{}
	}, func() rmreq.UpdateRequest[rmreq.PenilaianLanjutanResikoJatuhDewasaData] {
		return &rmreq.PenilaianLanjutanResikoJatuhDewasaUpdate{}
	}))
	reg("penilaian-lanjutan-resiko-jatuh-geriatri", rmctrl.NewController(rmaction.FormPenilaianLanjutanResikoJatuhGeriatri, func() rmreq.StoreRequest[rmreq.PenilaianLanjutanResikoJatuhGeriatriData] {
		return &rmreq.PenilaianLanjutanResikoJatuhGeriatriStore{}
	}, func() rmreq.UpdateRequest[rmreq.PenilaianLanjutanResikoJatuhGeriatriData] {
		return &rmreq.PenilaianLanjutanResikoJatuhGeriatriUpdate{}
	}))
	reg("penilaian-lanjutan-resiko-jatuh-lansia", rmctrl.NewController(rmaction.FormPenilaianLanjutanResikoJatuhLansia, func() rmreq.StoreRequest[rmreq.PenilaianLanjutanResikoJatuhLansiaData] {
		return &rmreq.PenilaianLanjutanResikoJatuhLansiaStore{}
	}, func() rmreq.UpdateRequest[rmreq.PenilaianLanjutanResikoJatuhLansiaData] {
		return &rmreq.PenilaianLanjutanResikoJatuhLansiaUpdate{}
	}))
	reg("penilaian-lanjutan-resiko-jatuh-psikiatri", rmctrl.NewController(rmaction.FormPenilaianLanjutanResikoJatuhPsikiatri, func() rmreq.StoreRequest[rmreq.PenilaianLanjutanResikoJatuhPsikiatriData] {
		return &rmreq.PenilaianLanjutanResikoJatuhPsikiatriStore{}
	}, func() rmreq.UpdateRequest[rmreq.PenilaianLanjutanResikoJatuhPsikiatriData] {
		return &rmreq.PenilaianLanjutanResikoJatuhPsikiatriUpdate{}
	}))
	reg("penilaian-lanjutan-skrining-fungsional", rmctrl.NewController(rmaction.FormPenilaianLanjutanSkriningFungsional, func() rmreq.StoreRequest[rmreq.PenilaianLanjutanSkriningFungsionalData] {
		return &rmreq.PenilaianLanjutanSkriningFungsionalStore{}
	}, func() rmreq.UpdateRequest[rmreq.PenilaianLanjutanSkriningFungsionalData] {
		return &rmreq.PenilaianLanjutanSkriningFungsionalUpdate{}
	}))
	reg("penilaian-level-kecemasan-ranap-anak", rmctrl.NewController(rmaction.FormPenilaianLevelKecemasanRanapAnak, func() rmreq.StoreRequest[rmreq.PenilaianLevelKecemasanRanapAnakData] {
		return &rmreq.PenilaianLevelKecemasanRanapAnakStore{}
	}, func() rmreq.UpdateRequest[rmreq.PenilaianLevelKecemasanRanapAnakData] {
		return &rmreq.PenilaianLevelKecemasanRanapAnakUpdate{}
	}))
	reg("penilaian-pasien-imunitas-rendah", rmctrl.NewController(rmaction.FormPenilaianPasienImunitasRendah, func() rmreq.StoreRequest[rmreq.PenilaianPasienImunitasRendahData] {
		return &rmreq.PenilaianPasienImunitasRendahStore{}
	}, func() rmreq.UpdateRequest[rmreq.PenilaianPasienImunitasRendahData] {
		return &rmreq.PenilaianPasienImunitasRendahUpdate{}
	}))
	reg("penilaian-pasien-keracunan", rmctrl.NewController(rmaction.FormPenilaianPasienKeracunan, func() rmreq.StoreRequest[rmreq.PenilaianPasienKeracunanData] {
		return &rmreq.PenilaianPasienKeracunanStore{}
	}, func() rmreq.UpdateRequest[rmreq.PenilaianPasienKeracunanData] {
		return &rmreq.PenilaianPasienKeracunanUpdate{}
	}))
	reg("penilaian-pasien-penyakit-menular", rmctrl.NewController(rmaction.FormPenilaianPasienPenyakitMenular, func() rmreq.StoreRequest[rmreq.PenilaianPasienPenyakitMenularData] {
		return &rmreq.PenilaianPasienPenyakitMenularStore{}
	}, func() rmreq.UpdateRequest[rmreq.PenilaianPasienPenyakitMenularData] {
		return &rmreq.PenilaianPasienPenyakitMenularUpdate{}
	}))
	reg("penilaian-pasien-terminal", rmctrl.NewController(rmaction.FormPenilaianPasienTerminal, func() rmreq.StoreRequest[rmreq.PenilaianPasienTerminalData] {
		return &rmreq.PenilaianPasienTerminalStore{}
	}, func() rmreq.UpdateRequest[rmreq.PenilaianPasienTerminalData] {
		return &rmreq.PenilaianPasienTerminalUpdate{}
	}))
	reg("penilaian-pre-anestesi", rmctrl.NewController(rmaction.FormPenilaianPreAnestesi, func() rmreq.StoreRequest[rmreq.PenilaianPreAnestesiData] { return &rmreq.PenilaianPreAnestesiStore{} }, func() rmreq.UpdateRequest[rmreq.PenilaianPreAnestesiData] { return &rmreq.PenilaianPreAnestesiUpdate{} }))
	reg("penilaian-pre-induksi", rmctrl.NewController(rmaction.FormPenilaianPreInduksi, func() rmreq.StoreRequest[rmreq.PenilaianPreInduksiData] { return &rmreq.PenilaianPreInduksiStore{} }, func() rmreq.UpdateRequest[rmreq.PenilaianPreInduksiData] { return &rmreq.PenilaianPreInduksiUpdate{} }))
	reg("penilaian-pre-operasi", rmctrl.NewController(rmaction.FormPenilaianPreOperasi, func() rmreq.StoreRequest[rmreq.PenilaianPreOperasiData] { return &rmreq.PenilaianPreOperasiStore{} }, func() rmreq.UpdateRequest[rmreq.PenilaianPreOperasiData] { return &rmreq.PenilaianPreOperasiUpdate{} }))
	reg("penilaian-psikologi", rmctrl.NewController(rmaction.FormPenilaianPsikologi, func() rmreq.StoreRequest[rmreq.PenilaianPsikologiData] { return &rmreq.PenilaianPsikologiStore{} }, func() rmreq.UpdateRequest[rmreq.PenilaianPsikologiData] { return &rmreq.PenilaianPsikologiUpdate{} }))
	reg("penilaian-psikologi-klinis", rmctrl.NewController(rmaction.FormPenilaianPsikologiKlinis, func() rmreq.StoreRequest[rmreq.PenilaianPsikologiKlinisData] {
		return &rmreq.PenilaianPsikologiKlinisStore{}
	}, func() rmreq.UpdateRequest[rmreq.PenilaianPsikologiKlinisData] {
		return &rmreq.PenilaianPsikologiKlinisUpdate{}
	}))
	reg("penilaian-risiko-dekubitus", rmctrl.NewController(rmaction.FormPenilaianRisikoDekubitus, func() rmreq.StoreRequest[rmreq.PenilaianRisikoDekubitusData] {
		return &rmreq.PenilaianRisikoDekubitusStore{}
	}, func() rmreq.UpdateRequest[rmreq.PenilaianRisikoDekubitusData] {
		return &rmreq.PenilaianRisikoDekubitusUpdate{}
	}))
	reg("penilaian-risiko-jatuh-neonatus", rmctrl.NewController(rmaction.FormPenilaianRisikoJatuhNeonatus, func() rmreq.StoreRequest[rmreq.PenilaianRisikoJatuhNeonatusData] {
		return &rmreq.PenilaianRisikoJatuhNeonatusStore{}
	}, func() rmreq.UpdateRequest[rmreq.PenilaianRisikoJatuhNeonatusData] {
		return &rmreq.PenilaianRisikoJatuhNeonatusUpdate{}
	}))
	reg("penilaian-tambahan-bunuh-diri", rmctrl.NewController(rmaction.FormPenilaianTambahanBunuhDiri, func() rmreq.StoreRequest[rmreq.PenilaianTambahanBunuhDiriData] {
		return &rmreq.PenilaianTambahanBunuhDiriStore{}
	}, func() rmreq.UpdateRequest[rmreq.PenilaianTambahanBunuhDiriData] {
		return &rmreq.PenilaianTambahanBunuhDiriUpdate{}
	}))
	reg("penilaian-tambahan-geriatri", rmctrl.NewController(rmaction.FormPenilaianTambahanGeriatri, func() rmreq.StoreRequest[rmreq.PenilaianTambahanGeriatriData] {
		return &rmreq.PenilaianTambahanGeriatriStore{}
	}, func() rmreq.UpdateRequest[rmreq.PenilaianTambahanGeriatriData] {
		return &rmreq.PenilaianTambahanGeriatriUpdate{}
	}))
	reg("penilaian-tambahan-beresiko-melarikan-diri", rmctrl.NewController(rmaction.FormPenilaianTambahanBeresikoMelarikanDiri, func() rmreq.StoreRequest[rmreq.PenilaianTambahanBeresikoMelarikanDiriData] {
		return &rmreq.PenilaianTambahanBeresikoMelarikanDiriStore{}
	}, func() rmreq.UpdateRequest[rmreq.PenilaianTambahanBeresikoMelarikanDiriData] {
		return &rmreq.PenilaianTambahanBeresikoMelarikanDiriUpdate{}
	}))
	reg("penilaian-tambahan-perilaku-kekerasan", rmctrl.NewController(rmaction.FormPenilaianTambahanPerilakuKekerasan, func() rmreq.StoreRequest[rmreq.PenilaianTambahanPerilakuKekerasanData] {
		return &rmreq.PenilaianTambahanPerilakuKekerasanStore{}
	}, func() rmreq.UpdateRequest[rmreq.PenilaianTambahanPerilakuKekerasanData] {
		return &rmreq.PenilaianTambahanPerilakuKekerasanUpdate{}
	}))
	reg("penilaian-terapi-wicara", rmctrl.NewController(rmaction.FormPenilaianTerapiWicara, func() rmreq.StoreRequest[rmreq.PenilaianTerapiWicaraData] { return &rmreq.PenilaianTerapiWicaraStore{} }, func() rmreq.UpdateRequest[rmreq.PenilaianTerapiWicaraData] {
		return &rmreq.PenilaianTerapiWicaraUpdate{}
	}))
	reg("penilaian-ulang-nyeri", rmctrl.NewController(rmaction.FormPenilaianUlangNyeri, func() rmreq.StoreRequest[rmreq.PenilaianUlangNyeriData] { return &rmreq.PenilaianUlangNyeriStore{} }, func() rmreq.UpdateRequest[rmreq.PenilaianUlangNyeriData] { return &rmreq.PenilaianUlangNyeriUpdate{} }))
	reg("perencanaan-pemulangan", rmctrl.NewController(rmaction.FormPerencanaanPemulangan, func() rmreq.StoreRequest[rmreq.PerencanaanPemulanganData] { return &rmreq.PerencanaanPemulanganStore{} }, func() rmreq.UpdateRequest[rmreq.PerencanaanPemulanganData] {
		return &rmreq.PerencanaanPemulanganUpdate{}
	}))
	reg("signin-sebelum-anestesi", rmctrl.NewController(rmaction.FormSigninSebelumAnestesi, func() rmreq.StoreRequest[rmreq.SigninSebelumAnestesiData] { return &rmreq.SigninSebelumAnestesiStore{} }, func() rmreq.UpdateRequest[rmreq.SigninSebelumAnestesiData] {
		return &rmreq.SigninSebelumAnestesiUpdate{}
	}))
	reg("signout-sebelum-menutup-luka", rmctrl.NewController(rmaction.FormSignoutSebelumMenutupLuka, func() rmreq.StoreRequest[rmreq.SignoutSebelumMenutupLukaData] {
		return &rmreq.SignoutSebelumMenutupLukaStore{}
	}, func() rmreq.UpdateRequest[rmreq.SignoutSebelumMenutupLukaData] {
		return &rmreq.SignoutSebelumMenutupLukaUpdate{}
	}))
	reg("skrining-adiksi-nikotin", rmctrl.NewController(rmaction.FormSkriningAdiksiNikotin, func() rmreq.StoreRequest[rmreq.SkriningAdiksiNikotinData] { return &rmreq.SkriningAdiksiNikotinStore{} }, func() rmreq.UpdateRequest[rmreq.SkriningAdiksiNikotinData] {
		return &rmreq.SkriningAdiksiNikotinUpdate{}
	}))
	reg("skrining-anemia", rmctrl.NewController(rmaction.FormSkriningAnemia, func() rmreq.StoreRequest[rmreq.SkriningAnemiaData] { return &rmreq.SkriningAnemiaStore{} }, func() rmreq.UpdateRequest[rmreq.SkriningAnemiaData] { return &rmreq.SkriningAnemiaUpdate{} }))
	reg("skrining-curb65", rmctrl.NewController(rmaction.FormSkriningCurb65, func() rmreq.StoreRequest[rmreq.SkriningCurb65Data] { return &rmreq.SkriningCurb65Store{} }, func() rmreq.UpdateRequest[rmreq.SkriningCurb65Data] { return &rmreq.SkriningCurb65Update{} }))
	reg("skrining-diabetes-melitus", rmctrl.NewController(rmaction.FormSkriningDiabetesMelitus, func() rmreq.StoreRequest[rmreq.SkriningDiabetesMelitusData] {
		return &rmreq.SkriningDiabetesMelitusStore{}
	}, func() rmreq.UpdateRequest[rmreq.SkriningDiabetesMelitusData] {
		return &rmreq.SkriningDiabetesMelitusUpdate{}
	}))
	reg("skrining-frailty-syndrome", rmctrl.NewController(rmaction.FormSkriningFrailtySyndrome, func() rmreq.StoreRequest[rmreq.SkriningFrailtySyndromeData] {
		return &rmreq.SkriningFrailtySyndromeStore{}
	}, func() rmreq.UpdateRequest[rmreq.SkriningFrailtySyndromeData] {
		return &rmreq.SkriningFrailtySyndromeUpdate{}
	}))
	reg("skrining-hipertensi", rmctrl.NewController(rmaction.FormSkriningHipertensi, func() rmreq.StoreRequest[rmreq.SkriningHipertensiData] { return &rmreq.SkriningHipertensiStore{} }, func() rmreq.UpdateRequest[rmreq.SkriningHipertensiData] { return &rmreq.SkriningHipertensiUpdate{} }))
	reg("skrining-indra-pendengaran", rmctrl.NewController(rmaction.FormSkriningIndraPendengaran, func() rmreq.StoreRequest[rmreq.SkriningIndraPendengaranData] {
		return &rmreq.SkriningIndraPendengaranStore{}
	}, func() rmreq.UpdateRequest[rmreq.SkriningIndraPendengaranData] {
		return &rmreq.SkriningIndraPendengaranUpdate{}
	}))
	reg("skrining-instrumen-acrs", rmctrl.NewController(rmaction.FormSkriningInstrumenAcrs, func() rmreq.StoreRequest[rmreq.SkriningInstrumenAcrsData] { return &rmreq.SkriningInstrumenAcrsStore{} }, func() rmreq.UpdateRequest[rmreq.SkriningInstrumenAcrsData] {
		return &rmreq.SkriningInstrumenAcrsUpdate{}
	}))
	reg("skrining-instrumen-amt", rmctrl.NewController(rmaction.FormSkriningInstrumenAmt, func() rmreq.StoreRequest[rmreq.SkriningInstrumenAmtData] { return &rmreq.SkriningInstrumenAmtStore{} }, func() rmreq.UpdateRequest[rmreq.SkriningInstrumenAmtData] { return &rmreq.SkriningInstrumenAmtUpdate{} }))
	reg("skrining-instrumen-esat", rmctrl.NewController(rmaction.FormSkriningInstrumenEsat, func() rmreq.StoreRequest[rmreq.SkriningInstrumenEsatData] { return &rmreq.SkriningInstrumenEsatStore{} }, func() rmreq.UpdateRequest[rmreq.SkriningInstrumenEsatData] {
		return &rmreq.SkriningInstrumenEsatUpdate{}
	}))
	reg("skrining-instrumen-sdq", rmctrl.NewController(rmaction.FormSkriningInstrumenSdq, func() rmreq.StoreRequest[rmreq.SkriningInstrumenSdqData] { return &rmreq.SkriningInstrumenSdqStore{} }, func() rmreq.UpdateRequest[rmreq.SkriningInstrumenSdqData] { return &rmreq.SkriningInstrumenSdqUpdate{} }))
	reg("skrining-kanker-kolorektal", rmctrl.NewController(rmaction.FormSkriningKankerKolorektal, func() rmreq.StoreRequest[rmreq.SkriningKankerKolorektalData] {
		return &rmreq.SkriningKankerKolorektalStore{}
	}, func() rmreq.UpdateRequest[rmreq.SkriningKankerKolorektalData] {
		return &rmreq.SkriningKankerKolorektalUpdate{}
	}))
	reg("skrining-kekerasan-pada-perempuan", rmctrl.NewController(rmaction.FormSkriningKekerasanPadaPerempuan, func() rmreq.StoreRequest[rmreq.SkriningKekerasanPadaPerempuanData] {
		return &rmreq.SkriningKekerasanPadaPerempuanStore{}
	}, func() rmreq.UpdateRequest[rmreq.SkriningKekerasanPadaPerempuanData] {
		return &rmreq.SkriningKekerasanPadaPerempuanUpdate{}
	}))
	reg("skrining-kesehatan-gigi-mulut-balita", rmctrl.NewController(rmaction.FormSkriningKesehatanGigiMulutBalita, func() rmreq.StoreRequest[rmreq.SkriningKesehatanGigiMulutBalitaData] {
		return &rmreq.SkriningKesehatanGigiMulutBalitaStore{}
	}, func() rmreq.UpdateRequest[rmreq.SkriningKesehatanGigiMulutBalitaData] {
		return &rmreq.SkriningKesehatanGigiMulutBalitaUpdate{}
	}))
	reg("skrining-kesehatan-gigi-mulut-dewasa", rmctrl.NewController(rmaction.FormSkriningKesehatanGigiMulutDewasa, func() rmreq.StoreRequest[rmreq.SkriningKesehatanGigiMulutDewasaData] {
		return &rmreq.SkriningKesehatanGigiMulutDewasaStore{}
	}, func() rmreq.UpdateRequest[rmreq.SkriningKesehatanGigiMulutDewasaData] {
		return &rmreq.SkriningKesehatanGigiMulutDewasaUpdate{}
	}))
	reg("skrining-kesehatan-gigi-mulut-lansia", rmctrl.NewController(rmaction.FormSkriningKesehatanGigiMulutLansia, func() rmreq.StoreRequest[rmreq.SkriningKesehatanGigiMulutLansiaData] {
		return &rmreq.SkriningKesehatanGigiMulutLansiaStore{}
	}, func() rmreq.UpdateRequest[rmreq.SkriningKesehatanGigiMulutLansiaData] {
		return &rmreq.SkriningKesehatanGigiMulutLansiaUpdate{}
	}))
	reg("skrining-kesehatan-gigi-mulut-remaja", rmctrl.NewController(rmaction.FormSkriningKesehatanGigiMulutRemaja, func() rmreq.StoreRequest[rmreq.SkriningKesehatanGigiMulutRemajaData] {
		return &rmreq.SkriningKesehatanGigiMulutRemajaStore{}
	}, func() rmreq.UpdateRequest[rmreq.SkriningKesehatanGigiMulutRemajaData] {
		return &rmreq.SkriningKesehatanGigiMulutRemajaUpdate{}
	}))
	reg("skrining-kesehatan-penglihatan", rmctrl.NewController(rmaction.FormSkriningKesehatanPenglihatan, func() rmreq.StoreRequest[rmreq.SkriningKesehatanPenglihatanData] {
		return &rmreq.SkriningKesehatanPenglihatanStore{}
	}, func() rmreq.UpdateRequest[rmreq.SkriningKesehatanPenglihatanData] {
		return &rmreq.SkriningKesehatanPenglihatanUpdate{}
	}))
	reg("mpp-skrining", rmctrl.NewController(rmaction.FormMppSkrining, func() rmreq.StoreRequest[rmreq.MppSkriningData] { return &rmreq.MppSkriningStore{} }, func() rmreq.UpdateRequest[rmreq.MppSkriningData] { return &rmreq.MppSkriningUpdate{} }))
	reg("mpp-evaluasi", rmctrl.NewController(rmaction.FormMppEvaluasi, func() rmreq.StoreRequest[rmreq.MppEvaluasiData] { return &rmreq.MppEvaluasiStore{} }, func() rmreq.UpdateRequest[rmreq.MppEvaluasiData] { return &rmreq.MppEvaluasiUpdate{} }))
	reg("mpp-evaluasi-catatan", rmctrl.NewController(rmaction.FormMppEvaluasiCatatan, func() rmreq.StoreRequest[rmreq.MppEvaluasiCatatanData] { return &rmreq.MppEvaluasiCatatanStore{} }, func() rmreq.UpdateRequest[rmreq.MppEvaluasiCatatanData] { return &rmreq.MppEvaluasiCatatanUpdate{} }))
	reg("skrining-perilaku-merokok-sekolah-remaja", rmctrl.NewController(rmaction.FormSkriningPerilakuMerokokSekolahRemaja, func() rmreq.StoreRequest[rmreq.SkriningPerilakuMerokokSekolahRemajaData] {
		return &rmreq.SkriningPerilakuMerokokSekolahRemajaStore{}
	}, func() rmreq.UpdateRequest[rmreq.SkriningPerilakuMerokokSekolahRemajaData] {
		return &rmreq.SkriningPerilakuMerokokSekolahRemajaUpdate{}
	}))
	reg("skrining-nutrisi-anak", rmctrl.NewController(rmaction.FormSkriningNutrisiAnak, func() rmreq.StoreRequest[rmreq.SkriningNutrisiAnakData] { return &rmreq.SkriningNutrisiAnakStore{} }, func() rmreq.UpdateRequest[rmreq.SkriningNutrisiAnakData] { return &rmreq.SkriningNutrisiAnakUpdate{} }))
	reg("skrining-nutrisi-dewasa", rmctrl.NewController(rmaction.FormSkriningNutrisiDewasa, func() rmreq.StoreRequest[rmreq.SkriningNutrisiDewasaData] { return &rmreq.SkriningNutrisiDewasaStore{} }, func() rmreq.UpdateRequest[rmreq.SkriningNutrisiDewasaData] {
		return &rmreq.SkriningNutrisiDewasaUpdate{}
	}))
	reg("skrining-nutrisi-lansia", rmctrl.NewController(rmaction.FormSkriningNutrisiLansia, func() rmreq.StoreRequest[rmreq.SkriningNutrisiLansiaData] { return &rmreq.SkriningNutrisiLansiaStore{} }, func() rmreq.UpdateRequest[rmreq.SkriningNutrisiLansiaData] {
		return &rmreq.SkriningNutrisiLansiaUpdate{}
	}))
	reg("skrining-obesitas", rmctrl.NewController(rmaction.FormSkriningObesitas, func() rmreq.StoreRequest[rmreq.SkriningObesitasData] { return &rmreq.SkriningObesitasStore{} }, func() rmreq.UpdateRequest[rmreq.SkriningObesitasData] { return &rmreq.SkriningObesitasUpdate{} }))
	reg("skrining-puma", rmctrl.NewController(rmaction.FormSkriningPuma, func() rmreq.StoreRequest[rmreq.SkriningPumaData] { return &rmreq.SkriningPumaStore{} }, func() rmreq.UpdateRequest[rmreq.SkriningPumaData] { return &rmreq.SkriningPumaUpdate{} }))
	reg("skrining-pneumonia-severity-index", rmctrl.NewController(rmaction.FormSkriningPneumoniaSeverityIndex, func() rmreq.StoreRequest[rmreq.SkriningPneumoniaSeverityIndexData] {
		return &rmreq.SkriningPneumoniaSeverityIndexStore{}
	}, func() rmreq.UpdateRequest[rmreq.SkriningPneumoniaSeverityIndexData] {
		return &rmreq.SkriningPneumoniaSeverityIndexUpdate{}
	}))
	reg("skrining-risiko-kanker-paru", rmctrl.NewController(rmaction.FormSkriningRisikoKankerParu, func() rmreq.StoreRequest[rmreq.SkriningRisikoKankerParuData] {
		return &rmreq.SkriningRisikoKankerParuStore{}
	}, func() rmreq.UpdateRequest[rmreq.SkriningRisikoKankerParuData] {
		return &rmreq.SkriningRisikoKankerParuUpdate{}
	}))
	reg("skrining-risiko-kanker-payudara", rmctrl.NewController(rmaction.FormSkriningRisikoKankerPayudara, func() rmreq.StoreRequest[rmreq.SkriningRisikoKankerPayudaraData] {
		return &rmreq.SkriningRisikoKankerPayudaraStore{}
	}, func() rmreq.UpdateRequest[rmreq.SkriningRisikoKankerPayudaraData] {
		return &rmreq.SkriningRisikoKankerPayudaraUpdate{}
	}))
	reg("skrining-risiko-kanker-serviks", rmctrl.NewController(rmaction.FormSkriningRisikoKankerServiks, func() rmreq.StoreRequest[rmreq.SkriningRisikoKankerServiksData] {
		return &rmreq.SkriningRisikoKankerServiksStore{}
	}, func() rmreq.UpdateRequest[rmreq.SkriningRisikoKankerServiksData] {
		return &rmreq.SkriningRisikoKankerServiksUpdate{}
	}))
	reg("skrining-instrumen-srq", rmctrl.NewController(rmaction.FormSkriningInstrumenSrq, func() rmreq.StoreRequest[rmreq.SkriningInstrumenSrqData] { return &rmreq.SkriningInstrumenSrqStore{} }, func() rmreq.UpdateRequest[rmreq.SkriningInstrumenSrqData] { return &rmreq.SkriningInstrumenSrqUpdate{} }))
	reg("skrining-tbc", rmctrl.NewController(rmaction.FormSkriningTbc, func() rmreq.StoreRequest[rmreq.SkriningTbcData] { return &rmreq.SkriningTbcStore{} }, func() rmreq.UpdateRequest[rmreq.SkriningTbcData] { return &rmreq.SkriningTbcUpdate{} }))
	reg("skrining-tolac", rmctrl.NewController(rmaction.FormSkriningTolac, func() rmreq.StoreRequest[rmreq.SkriningTolacData] { return &rmreq.SkriningTolacStore{} }, func() rmreq.UpdateRequest[rmreq.SkriningTolacData] { return &rmreq.SkriningTolacUpdate{} }))
	reg("skrining-thalassemia", rmctrl.NewController(rmaction.FormSkriningThalassemia, func() rmreq.StoreRequest[rmreq.SkriningThalassemiaData] { return &rmreq.SkriningThalassemiaStore{} }, func() rmreq.UpdateRequest[rmreq.SkriningThalassemiaData] { return &rmreq.SkriningThalassemiaUpdate{} }))
	reg("timeout-sebelum-insisi", rmctrl.NewController(rmaction.FormTimeoutSebelumInsisi, func() rmreq.StoreRequest[rmreq.TimeoutSebelumInsisiData] { return &rmreq.TimeoutSebelumInsisiStore{} }, func() rmreq.UpdateRequest[rmreq.TimeoutSebelumInsisiData] { return &rmreq.TimeoutSebelumInsisiUpdate{} }))
	reg("transfer-pasien-antar-ruang", rmctrl.NewController(rmaction.FormTransferPasienAntarRuang, func() rmreq.StoreRequest[rmreq.TransferPasienAntarRuangData] {
		return &rmreq.TransferPasienAntarRuangStore{}
	}, func() rmreq.UpdateRequest[rmreq.TransferPasienAntarRuangData] {
		return &rmreq.TransferPasienAntarRuangUpdate{}
	}))
	reg("uji-fungsi-kfr", rmctrl.NewController(rmaction.FormUjiFungsiKfr, func() rmreq.StoreRequest[rmreq.UjiFungsiKfrData] { return &rmreq.UjiFungsiKfrStore{} }, func() rmreq.UpdateRequest[rmreq.UjiFungsiKfrData] { return &rmreq.UjiFungsiKfrUpdate{} }))
}
